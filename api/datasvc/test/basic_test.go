package test

import (
	"fmt"
	"path"
	"sort"
	"strconv"
	"time"

	check "github.com/erikh/check"
	"github.com/golang/mock/gomock"
	"github.com/tinyci/ci-agents/config"
	"github.com/tinyci/ci-agents/mocks/github"
	"github.com/tinyci/ci-agents/model"
	"github.com/tinyci/ci-agents/testutil"
	"golang.org/x/oauth2"
)

func (ds *datasvcSuite) TestBasicUser(c *check.C) {
	username := testutil.RandString(8)
	resp, err := ds.client.MakeUser(username)
	c.Assert(err, check.IsNil)
	c.Assert(resp.ID, check.Not(check.Equals), int64(0))
	c.Assert(resp.Username, check.Equals, username)

	user, err := ds.client.Client().GetUser(username)
	c.Assert(err, check.IsNil)
	c.Assert(resp, check.DeepEquals, user)

	user.Token = &oauth2.Token{AccessToken: "this is in this test"}
	c.Assert(ds.client.Client().PatchUser(user), check.IsNil)
	user, err = ds.client.Client().GetUser(username)
	c.Assert(err, check.IsNil)
	c.Assert(user.Token, check.NotNil)
	c.Assert(user.Token.AccessToken, check.Equals, "this is in this test")

	user.Token = testutil.DummyToken
	c.Assert(ds.client.Client().PatchUser(user), check.IsNil)

	now := time.Now() // should not be able to update this field.
	user.LastScannedRepos = &now
	c.Assert(ds.client.Client().PatchUser(user), check.IsNil)
	user, err = ds.client.Client().GetUser(username)
	c.Assert(err, check.IsNil)
	c.Assert(user.LastScannedRepos, check.Not(check.Equals), &now)

	user.Username = "notcool"
	c.Assert(ds.client.Client().PatchUser(user), check.NotNil)

	users := map[string]*model.User{ // preseed with already written record
		username: resp,
	}

	for i := 0; i < 10; i++ {
		username := testutil.RandString(8)
		resp, err := ds.client.MakeUser(username)
		c.Assert(err, check.IsNil)
		users[username] = resp
	}

	usersResp, err := ds.client.Client().ListUsers()
	c.Assert(err, check.IsNil)

	for _, item := range usersResp {
		m, ok := users[item.Username]
		c.Assert(ok, check.Equals, true, check.Commentf("%s", item.Username))
		c.Assert(m, check.DeepEquals, item, check.Commentf("%s", item.Username))
	}
}

func (ds *datasvcSuite) TestUserLoginToken(c *check.C) {
	username := testutil.RandString(8)
	resp, err := ds.client.MakeUser(username)
	c.Assert(err, check.IsNil)
	c.Assert(resp.ID, check.Not(check.Equals), int64(0))
	c.Assert(resp.Username, check.Equals, username)

	_, err = ds.client.Client().GetToken("quux")
	c.Assert(err, check.NotNil)

	accessToken, err := ds.client.Client().GetToken(username)
	c.Assert(err, check.IsNil)

	_, err = ds.client.Client().GetToken(username)
	c.Assert(err, check.NotNil)

	c.Assert(ds.client.Client().DeleteToken(username), check.IsNil)

	accessToken2, err := ds.client.Client().GetToken(username)
	c.Assert(err, check.IsNil)

	c.Assert(accessToken, check.Not(check.Equals), accessToken2)

	_, err = ds.client.Client().ValidateToken(accessToken)
	c.Assert(err, check.NotNil)
	u, err := ds.client.Client().ValidateToken(accessToken2)
	c.Assert(err, check.IsNil)
	c.Assert(u.Username, check.Equals, username)
}

func (ds *datasvcSuite) TestUserErrors(c *check.C) {
	username := testutil.RandString(8)
	_, err := ds.client.MakeUser(username)
	c.Assert(err, check.IsNil)
	messages := []string{}

	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("%d", i)
		messages = append(messages, msg)

		c.Assert(
			ds.client.Client().AddError(msg, username),
			check.IsNil,
		)
	}

	resp, err := ds.client.Client().GetErrors(username)
	c.Assert(err, check.IsNil)
	c.Assert(len(resp), check.Equals, 10)
	respStr := []string{}

	for _, item := range resp {
		respStr = append(respStr, string(item.Error))
	}

	sort.Strings(respStr)
	c.Assert(messages, check.DeepEquals, respStr)
}

func (ds *datasvcSuite) TestRepositories(c *check.C) {
	username := testutil.RandString(8)
	_, err := ds.client.MakeUser(username)
	c.Assert(err, check.IsNil)

	repos := []string{}

	for i := 0; i < 10; i++ {
		owner := testutil.RandString(8) + "_" + strconv.Itoa(i)
		repo := testutil.RandString(8)
		fullRepo := path.Join(owner, repo)
		c.Assert(ds.client.MakeRepo(fullRepo, username, i%2 == 0), check.IsNil)
		repos = append(repos, fullRepo)
	}

	all, err := ds.client.Client().AllRepositories(username, "")
	c.Assert(err, check.IsNil)
	c.Assert(len(all), check.Equals, 10)

	all, err = ds.client.Client().AllRepositories(username, "_2")
	c.Assert(err, check.IsNil)
	c.Assert(len(all), check.Equals, 1)

	private, err := ds.client.Client().PrivateRepositories(username, "")
	c.Assert(err, check.IsNil)
	c.Assert(len(private), check.Equals, 5)

	private, err = ds.client.Client().PrivateRepositories(username, "_2")
	c.Assert(err, check.IsNil)
	c.Assert(len(private), check.Equals, 1)

	public, err := ds.client.Client().PublicRepositories("")
	c.Assert(err, check.IsNil)
	c.Assert(len(public), check.Equals, 5)

	public, err = ds.client.Client().PublicRepositories("_3")
	c.Assert(err, check.IsNil)
	c.Assert(len(public), check.Equals, 1)

	repo, err := ds.client.Client().GetRepository(repos[0])
	c.Assert(err, check.IsNil)
	c.Assert(repo.Name, check.Equals, repos[0])
	c.Assert(repo.Disabled, check.Equals, true)
	c.Assert(repo.HookSecret, check.Equals, "")

	c.Assert(ds.client.Client().EnableRepository(username, repos[0]), check.IsNil)
	repo, err = ds.client.Client().GetRepository(repos[0])
	c.Assert(err, check.IsNil)
	c.Assert(repo.HookSecret, check.Not(check.Equals), "")
	c.Assert(repo.Disabled, check.Equals, false)

	c.Assert(ds.client.Client().DisableRepository(username, repos[0]), check.IsNil)
	repo, err = ds.client.Client().GetRepository(repos[0])
	c.Assert(err, check.IsNil)
	c.Assert(repo.Disabled, check.Equals, true)
}

func (ds *datasvcSuite) TestSubscriptions(c *check.C) {
	username := testutil.RandString(8)
	_, err := ds.client.MakeUser(username)
	c.Assert(err, check.IsNil)

	repos := []string{}

	for i := 0; i < 10; i++ {
		owner := testutil.RandString(8) + "_" + strconv.Itoa(i)
		repo := testutil.RandString(8)
		fullRepo := path.Join(owner, repo)
		c.Assert(ds.client.MakeRepo(fullRepo, username, false), check.IsNil)
		repos = append(repos, fullRepo)
	}

	username2 := testutil.RandString(8)
	_, err = ds.client.MakeUser(username2)
	c.Assert(err, check.IsNil)

	for _, repo := range repos {
		c.Assert(
			ds.client.Client().AddSubscription(username2, repo),
			check.IsNil,
		)
	}

	subs, err := ds.client.Client().ListSubscriptions(username2, "")
	c.Assert(err, check.IsNil)
	c.Assert(len(subs), check.Equals, 10)

	subs, err = ds.client.Client().ListSubscriptions(username2, "_1")
	c.Assert(err, check.IsNil)
	c.Assert(len(subs), check.Equals, 1)

	c.Assert(ds.client.MakeRepo("erikh/private", username, true), check.IsNil)
	c.Assert(
		ds.client.Client().AddSubscription(username2, "erikh/private"),
		check.NotNil,
	)
}

func (ds *datasvcSuite) TestRuns(c *check.C) {
	config.DefaultGithubClient = github.NewMockClient(gomock.NewController(c))
	now := time.Now()
	qis := []*model.QueueItem{}
	for i := 0; i < 1000; i++ {
		qi, err := ds.client.MakeQueueItem()
		c.Assert(err, check.IsNil)
		qis = append(qis, qi)
	}

	fmt.Printf("Filling queue took %v\n", time.Since(now))

	count, err := ds.client.Client().RunCount("", "")
	c.Assert(err, check.IsNil)
	c.Assert(count, check.Equals, int64(1000))

	count, err = ds.client.Client().RunCount(qis[0].Run.Task.Ref.Repository.Name, "")
	c.Assert(err, check.IsNil)
	c.Assert(count, check.Equals, int64(1))

	count, err = ds.client.Client().RunCount(qis[0].Run.Task.Ref.Repository.Name, "foo")
	c.Assert(err, check.IsNil)
	c.Assert(count, check.Equals, int64(0))

	count, err = ds.client.Client().RunCount(qis[0].Run.Task.Ref.Repository.Name, qis[0].Run.Task.Ref.SHA)
	c.Assert(err, check.IsNil)
	c.Assert(count, check.Equals, int64(1))

	for i := 0; i < 10; i++ {
		runs, err := ds.client.Client().ListRuns("", "", int64(i), 100)
		c.Assert(err, check.IsNil)
		c.Assert(len(runs), check.Equals, 100, check.Commentf("Loop: %d", i))
	}

	runs, err := ds.client.Client().ListRuns("", "", 10, 100)
	c.Assert(err, check.IsNil)
	c.Assert(len(runs), check.Equals, 0)

	_, err = ds.client.Client().ListRuns("", "", -1, 100)
	c.Assert(err, check.NotNil)
}

func (ds *datasvcSuite) TestQueue(c *check.C) {
	config.DefaultGithubClient = github.NewMockClient(gomock.NewController(c))
	now := time.Now()
	for i := 0; i < 1000; i++ {
		_, err := ds.client.MakeQueueItem()
		c.Assert(err, check.IsNil)
	}

	fmt.Printf("Filling queue took %v\n", time.Since(now))

	for i := 0; i < 1000; i++ {
		qi, err := ds.client.Client().NextQueueItem("default", "hi")
		c.Assert(err, check.IsNil)
		c.Assert(qi.Running, check.Equals, true)
		c.Assert(qi.RunningOn, check.NotNil)
		c.Assert(*qi.RunningOn, check.Equals, "hi")
		c.Assert(qi.StartedAt, check.NotNil)
		c.Assert(qi.Run.StartedAt, check.NotNil)
	}

	_, err := ds.client.Client().NextQueueItem("default", "hi")
	c.Assert(err, check.NotNil)
}

func (ds *datasvcSuite) TestOAuth(c *check.C) {
	c.Assert(ds.client.Client().OAuthRegisterState("asdf"), check.IsNil)
	c.Assert(ds.client.Client().OAuthValidateState("asdf"), check.IsNil)
	c.Assert(ds.client.Client().OAuthValidateState("asdf2"), check.NotNil)
}

func (ds *datasvcSuite) TestRef(c *check.C) {
	username := testutil.RandString(8)
	_, err := ds.client.MakeUser(username)
	c.Assert(err, check.IsNil)

	ownerName, repoName := testutil.RandString(8), testutil.RandString(8)

	c.Assert(ds.client.MakeRepo(path.Join(ownerName, repoName), username, false), check.IsNil)

	repo, err := ds.client.Client().GetRepository(path.Join(ownerName, repoName))
	c.Assert(err, check.IsNil)

	id, err := ds.client.Client().PutRef(&model.Ref{
		Repository: repo,
		RefName:    "heads/hi",
		SHA:        "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	})
	c.Assert(err, check.IsNil)
	c.Assert(id, check.Not(check.Equals), int64(0))

	ref, err := ds.client.Client().GetRefByNameAndSHA(path.Join(ownerName, repoName), "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	c.Assert(err, check.IsNil)
	c.Assert(ref.ID, check.Equals, id)

	_, err = ds.client.Client().GetRefByNameAndSHA(path.Join(ownerName, repoName), "baaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	c.Assert(err, check.NotNil)

	_, err = ds.client.Client().GetRefByNameAndSHA(path.Join(testutil.RandString(8), repoName), "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	c.Assert(err, check.NotNil)

	_, err = ds.client.Client().GetRefByNameAndSHA(path.Join(ownerName, testutil.RandString(8)), "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	c.Assert(err, check.NotNil)
}
