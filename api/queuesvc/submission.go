package queuesvc

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/google/go-github/github"
	"github.com/tinyci/ci-agents/ci-gen/grpc/handler"
	gh "github.com/tinyci/ci-agents/clients/github"
	"github.com/tinyci/ci-agents/clients/log"
	"github.com/tinyci/ci-agents/errors"
	"github.com/tinyci/ci-agents/model"
	"github.com/tinyci/ci-agents/types"
)

const (
	defaultMainBranch  = "heads/master"
	repoConfigFilename = "tinyci.yml"
	taskConfigFilename = "task.yml"
)

type queueItems []*model.QueueItem

func (qi queueItems) Swap(i, j int) {
	qi[i], qi[j] = qi[j], qi[i]
}

func (qi queueItems) Less(i, j int) bool {
	return qi[i].Run.Name < qi[j].Run.Name
}

func (qi queueItems) Len() int {
	return len(qi)
}

type submissionProcessor struct {
	handler    *handler.H
	logger     *log.SubLogger
	submission *types.Submission

	parent *model.Repository
	fork   *model.Repository

	ghClient gh.Client

	repoConfig *types.RepoConfig

	root string
}

func getLogger(sub *types.Submission, h *handler.H) *log.SubLogger {
	if sub != nil {
		return h.Clients.Log.WithFields(log.FieldMap{
			"parent":       sub.Parent,
			"fork":         sub.Fork,
			"head":         sub.HeadSHA,
			"base":         sub.BaseSHA,
			"manual":       fmt.Sprintf("%v", sub.Manual),
			"submitted_by": sub.SubmittedBy,
			"all":          fmt.Sprintf("%v", sub.All),
		})
	}
	return h.Clients.Log
}

func (qs *QueueServer) newSubmissionProcessor(sub *types.Submission) *submissionProcessor {
	return &submissionProcessor{handler: qs.H, logger: getLogger(sub, qs.H), submission: sub}
}

func (sp *submissionProcessor) cleanup() *errors.Error {
	if sp.root != "" {
		return errors.New(os.RemoveAll(sp.root))
	}

	return nil
}

func (sp *submissionProcessor) process(ctx context.Context) ([]*model.QueueItem, *errors.Error) {
	if err := sp.submission.Validate(); err != nil {
		return nil, err
	}

	var eErr error

	sp.root, eErr = ioutil.TempDir("", "")
	if eErr != nil {
		return nil, errors.New(eErr).Wrap("while creating directory for git clone")
	}

	var err *errors.Error

	if sp.submission.Manual {
		user, err := sp.handler.Clients.Data.GetUser(ctx, sp.submission.SubmittedBy)
		if err != nil {
			return nil, err.Wrap("could not find user account")
		}

		sp.ghClient = sp.handler.OAuth.GithubClient(user.Token)
	} else {
		sp.parent, err = sp.handler.Clients.Data.GetRepository(ctx, sp.submission.Parent)
		if err != nil {
			return nil, err.Wrap("while retrieving parent repository")
		}

		sp.ghClient = sp.handler.OAuth.GithubClient(sp.parent.Owner.Token)
	}

forkRetry:
	sp.fork, err = sp.handler.Clients.Data.GetRepository(ctx, sp.submission.Fork)
	if err != nil {
		if err.Contains(errors.ErrNotFound) {
			ghFork, err := sp.ghClient.GetRepository(ctx, sp.submission.Fork)
			if err != nil {
				return nil, err.Wrap("while trying to fetch the fork to add")
			}

			if err := sp.handler.Clients.Data.PutRepositories(ctx, sp.submission.Fork, []*github.Repository{ghFork}, true); err != nil {
				return nil, err.Wrap("auto-creating fork repository")
			}

			goto forkRetry
		} else {
			return nil, err.Wrap("while retrieving fork repository")
		}
	}

	if sp.submission.Manual {
		sp.submission.BaseSHA = sp.submission.HeadSHA
		sp.parent = sp.fork
	}

	repo, eErr := git.PlainCloneContext(ctx, sp.root, false, &git.CloneOptions{
		URL: sp.parent.Github.GetCloneURL(),
		Auth: &http.BasicAuth{
			Username: "x", // anything except an empty string
			Password: sp.parent.Owner.Token.Token,
		},
		RecurseSubmodules: 5, // FIXME make this configurable
	})
	if eErr != nil {
		return nil, errors.New(eErr)
	}

	forkRemote, eErr := repo.CreateRemote(&config.RemoteConfig{
		Name: "fork",
		URLs: []string{sp.fork.Github.GetCloneURL()},
	})

	if eErr != nil {
		return nil, errors.New(eErr)
	}

	eErr = forkRemote.FetchContext(ctx, &git.FetchOptions{
		Auth: &http.BasicAuth{
			Username: "x", // anything except an empty string
			Password: sp.parent.Owner.Token.Token,
		},
	})
	if eErr != nil {
		return nil, errors.New(eErr)
	}

	baseref, eErr := repo.Reference(plumbing.ReferenceName("refs/"+sp.submission.BaseSHA), false)
	if eErr != nil {
		return nil, errors.New(eErr).Wrapf("locating base commit @ %q", sp.submission.BaseSHA)
	}

	baseObj, eErr := repo.CommitObject(baseref.Hash())
	if eErr != nil {
		return nil, errors.New(eErr).Wrapf("locating base commit obj @ %q", sp.submission.BaseSHA)
	}

	baseTree, eErr := repo.TreeObject(baseObj.TreeHash)
	if eErr != nil {
		return nil, errors.New(eErr).Wrapf("locating base ref tree @ %q", sp.submission.BaseSHA)
	}

	headref, eErr := repo.Reference(plumbing.ReferenceName("refs/"+sp.submission.HeadSHA), false)
	if eErr != nil {
		return nil, errors.New(eErr).Wrapf("locating head commit @ %q", sp.submission.HeadSHA)
	}

	headObj, eErr := repo.CommitObject(headref.Hash())
	if eErr != nil {
		return nil, errors.New(eErr).Wrapf("locating head commit obj @ %q", sp.submission.BaseSHA)
	}

	headTree, eErr := repo.TreeObject(headObj.TreeHash)
	if eErr != nil {
		return nil, errors.New(eErr).Wrapf("locating head ref tree @ %q", sp.submission.HeadSHA)
	}

	changes, eErr := headTree.DiffContext(ctx, baseTree)
	if eErr != nil {
		return nil, errors.New(eErr).Wrap("generating diff")
	}

	patch, eErr := changes.PatchContext(ctx)
	if eErr != nil {
		return nil, errors.New(eErr).Wrap("generating patch")
	}

	dirs := map[string]struct{}{}

	for _, filePatch := range patch.FilePatches() {
		to, from := filePatch.Files()
		dirs[filepath.Dir(to.Path())] = struct{}{}
		dirs[filepath.Dir(from.Path())] = struct{}{}
	}

	wt, eErr := repo.Worktree()
	if eErr != nil {
		return nil, errors.New(eErr)
	}

	if err := wt.Checkout(&git.CheckoutOptions{Hash: headObj.Hash}); err != nil {
		return nil, errors.New(err)
	}

	taskymls := map[string]struct{}{}

	eErr = filepath.Walk(sp.root, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// try to do this in a single iteration
		if filepath.Base(p) == taskConfigFilename {
			rel, err := filepath.Rel(wt.Filesystem.Root(), p)
			if err != nil {
				return err
			}

			if strings.HasPrefix(rel, "..") {
				return errors.New("oh my how did you get here? please report a bug in this relative pathing issue")
			}

			taskymls[filepath.Dir(p)] = struct{}{}
		}

		return nil
	})
	if eErr != nil {
		return nil, errors.New(eErr)
	}

	selected := map[string]struct{}{}

	for dir := range dirs {
		d := dir

		// "." means the string is empty according to the docs
		for d != "." {
			if _, ok := taskymls[d]; ok {
				selected[d] = struct{}{}
				break
			}
			d = filepath.Dir(d)
		}
	}

	if err := sp.getRepoConfig(); err != nil {
		return nil, err
	}

	// FIXME do stuff with selected
	subRecord, err := sp.makeSub(ctx)
	if err != nil {
		return nil, err.Wrap("while writing submission")
	}

	tasks, taskDirs, err := sp.makeTaskDirs(ctx, selected, subRecord)
	if err != nil {
		return nil, err.Wrap("while configuring tasks")
	}

	qis := queueItems{}

	for _, dir := range taskDirs {
		qi, err := sp.generateQueueItems(ctx, dir, tasks[dir])
		if err != nil {
			return nil, err
		}

		qis = append(qis, qi...)
	}

	sort.Sort(qis)
	return qis, errors.New("unimplemented")
}

func (sp *submissionProcessor) manageRef(ctx context.Context, repo *model.Repository, sha string) (*model.Ref, *errors.Error) {
	ref, err := sp.handler.Clients.Data.GetRefByNameAndSHA(ctx, repo.Name, sha)
	if err != nil {
		if err.Contains(errors.ErrNotFound) {
			refs, err := sp.ghClient.GetRefs(ctx, repo.Name, sha)
			if err != nil {
				return nil, err.Wrap("while retrieving ref in submission")
			}

			var refName string

			if len(refs) > 0 {
				sort.Strings(refs)
				refName = refs[0]
			} else {
				refName = sha
			}

			ref = &model.Ref{Repository: repo, RefName: refName, SHA: sha}
			id, err := sp.handler.Clients.Data.PutRef(ctx, ref)
			if err != nil {
				return nil, err.Wrap("while adding new ref to collection")
			}

			ref.ID = id
			return ref, nil
		}

		return nil, err
	}

	return ref, nil
}

func (sp *submissionProcessor) makeSub(ctx context.Context) (*model.Submission, *errors.Error) {
	forkRef, err := sp.manageRef(ctx, sp.fork, sp.submission.HeadSHA)
	if err != nil {
		return nil, err.Wrap("while gathering fork ref")
	}

	var parentRef *model.Ref

	if sp.submission.Manual {
		parentRef = forkRef
	} else {
		var err *errors.Error
		parentRef, err = sp.manageRef(ctx, sp.parent, sp.submission.BaseSHA)
		if err != nil {
			return nil, err.Wrap("while gathering parent ref")
		}
	}

	return sp.handler.Clients.Data.PutSubmission(ctx, &model.Submission{
		TicketID: sp.submission.TicketID,
		HeadRef:  forkRef,
		BaseRef:  parentRef,
	})
}

func (sp *submissionProcessor) getRepoConfig() *errors.Error {
	content, eErr := ioutil.ReadFile(filepath.Join(sp.root, repoConfigFilename))
	if eErr != nil {
		return errors.New(eErr).Wrap("while reading tinyCI in-repository configuration")
	}

	rc, err := types.NewRepoConfig(content)
	if err != nil {
		return err.Wrap("while parsing tinyCI in-repository configuration")
	}

	sp.repoConfig = rc
	return nil
}

func (sp *submissionProcessor) makeTask(ctx context.Context, dir string, subRecord *model.Submission) (*model.Task, *errors.Error) {
	content, eErr := ioutil.ReadFile(filepath.Join(sp.root, dir, taskConfigFilename))
	if eErr != nil {
		return nil, errors.New(eErr).Wrapf("while reading task configuration in dir %q", dir)
	}

	ts, err := types.NewTaskSettings(content, false, sp.repoConfig)
	if err != nil {
		if sp.submission.TicketID != 0 {
			if cerr := sp.ghClient.CommentError(ctx, sp.parent.Name, sp.submission.TicketID, err.Wrap("tinyCI had an error processing your pull request")); cerr != nil {
				return nil, cerr.Wrap("attempting to alert the user about the error in their pull request")
			}
		}

		return nil, err.Wrapf("validating task settings for repo %q sha %q dir %q", sp.fork.Name, sp.submission.HeadSHA, dir)
	}

	return &model.Task{
		Path:         dir,
		TaskSettings: ts,
		CreatedAt:    time.Now(),
		Submission:   subRecord,
	}, nil
}

func (sp *submissionProcessor) makeTaskDirs(ctx context.Context, process map[string]struct{}, subRecord *model.Submission) (map[string]*model.Task, []string, *errors.Error) {
	tasks := map[string]*model.Task{}

	sp.logger.Info(ctx, "Computing task dirs")

	taskdirs := []string{}
	for dir := range process {
		taskdirs = append(taskdirs, dir)
	}

	for i := 0; i < len(taskdirs); i++ {
		task, err := sp.makeTask(ctx, taskdirs[i], subRecord)
		if err != nil {
			return nil, nil, err.Wrap("making task")
		}

		tasks[taskdirs[i]] = task

		for _, dir := range task.TaskSettings.Dependencies {
			if _, ok := process[dir]; !ok {
				process[dir] = struct{}{}
				taskdirs = append(taskdirs, dir)
			}
		}
	}

	sort.Strings(taskdirs)

	return tasks, taskdirs, nil
}

func (sp *submissionProcessor) generateQueueItems(ctx context.Context, dir string, task *model.Task) (queueItems, *errors.Error) {
	qis := queueItems{}

	task, err := sp.handler.Clients.Data.PutTask(ctx, task)
	if err != nil {
		return nil, err.Wrap("Could not insert task")
	}

	names := []string{}

	for name := range task.TaskSettings.Runs {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		qi, err := sp.makeRunQueue(ctx, name, dir, task)
		if err != nil {
			return nil, err.Wrap("constructing queue item")
		}
		qis = append(qis, qi)
	}

	return qis, nil
}

func (sp *submissionProcessor) makeRunQueue(ctx context.Context, name, dir string, task *model.Task) (*model.QueueItem, *errors.Error) {
	rs := task.TaskSettings.Runs[name]

	dirStr := dir

	if dir == "." || dir == "" {
		dirStr = "*root*"
	}

	run := &model.Run{
		Name:        strings.Join([]string{dirStr, name}, ":"),
		RunSettings: rs,
		Task:        task,
		CreatedAt:   time.Now(),
	}

	go sp.setPendingStatus(ctx, run)

	return &model.QueueItem{
		Run:       run,
		QueueName: run.RunSettings.Queue,
	}, nil
}

func (sp *submissionProcessor) setPendingStatus(ctx context.Context, run *model.Run) {
	parts := strings.SplitN(sp.parent.Name, "/", 2)
	if len(parts) != 2 {
		sp.logger.Error(ctx, errors.Errorf("invalid repo name %q", sp.parent.Name))
	}

	if err := sp.ghClient.PendingStatus(ctx, parts[0], parts[1], run.Name, sp.submission.HeadSHA, sp.handler.URL); err != nil {
		sp.logger.Error(ctx, err.Wrap("could not set pending status"))
	}
}
