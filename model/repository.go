package model

import (
	"encoding/base32"
	"encoding/json"
	"fmt"
	"strings"

	gh "github.com/google/go-github/github"
	"github.com/gorilla/securecookie"
	"github.com/jinzhu/gorm"
	"github.com/tinyci/ci-agents/errors"
	"github.com/tinyci/ci-agents/grpc/types"
	"github.com/tinyci/ci-agents/utils"
)

// RepositoryList conforms to the sort.Interface interface
type RepositoryList []*Repository

// Len computes the length of the list
func (rl RepositoryList) Len() int {
	return len(rl)
}

// Less determines the order of the list
func (rl RepositoryList) Less(i, j int) bool {
	return strings.Compare(rl[i].Name, rl[j].Name) < 0
}

func (rl RepositoryList) Swap(i, j int) {
	rl[j], rl[i] = rl[i], rl[j]
}

// ToProto converts the repository list to a protobuf representation
func (rl RepositoryList) ToProto() *types.RepositoryList {
	ret := &types.RepositoryList{}

	for _, repo := range rl {
		ret.List = append(ret.List, repo.ToProto())
	}

	return ret
}

// Repository is the encapsulation of a git repository.
type Repository struct {
	ID          int64  `gorm:"primary_key" json:"id"`
	Name        string `gorm:"unique" json:"name"`
	Private     bool   `json:"private"`
	Disabled    bool   `json:"disabled"`
	GithubJSON  []byte `gorm:"column:github" json:"-"`
	OwnerID     int64  `json:"-"`
	Owner       *User  `gorm:"association_autoupdate:false" json:"-"`
	AutoCreated bool   `json:"auto_created"`
	HookSecret  string

	Github *gh.Repository `json:"github"`
}

// NewRepositoryFromProto converts a proto repository to a model repository.
func NewRepositoryFromProto(r *types.Repository) (*Repository, *errors.Error) {
	github := &gh.Repository{}
	if err := json.Unmarshal(r.Github, github); err != nil {
		return nil, errors.New(err)
	}

	var owner *User

	if r.Owner != nil {
		var err *errors.Error
		owner, err = NewUserFromProto(r.Owner)
		if err != nil {
			return nil, err
		}
	}

	return &Repository{
		ID:          r.Id,
		Name:        r.Name,
		Private:     r.Private,
		Disabled:    r.Disabled,
		Owner:       owner,
		AutoCreated: r.AutoCreated,
		HookSecret:  r.HookSecret,
		Github:      github,
		GithubJSON:  r.Github,
	}, nil
}

// ToProto returns the protobuf representation of the repository
func (r *Repository) ToProto() *types.Repository {
	var owner *types.User

	if r.Owner != nil {
		owner = r.Owner.ToProto()
	}

	return &types.Repository{
		Id:          r.ID,
		Name:        r.Name,
		Private:     r.Private,
		Disabled:    r.Disabled,
		Owner:       owner,
		AutoCreated: r.AutoCreated,
		HookSecret:  r.HookSecret,
		Github:      r.GithubJSON,
	}
}

// OwnerRepo validates the owner/repo github path then returns each part.
func (r *Repository) OwnerRepo() (string, string, *errors.Error) {
	return utils.OwnerRepo(r.Name)
}

// GetRepositoryByNameForUser retrieves the repository by name if the user can
// see it (aka, if it's not private or if it's owned by them)
func (m *Model) GetRepositoryByNameForUser(name string, u *User) (*Repository, *errors.Error) {
	r := &Repository{}

	var id int64
	if u != nil {
		id = u.ID
	}

	return r, m.WrapError(m.Where("(owner_id = ? or not private) and name = ?", id, name).First(r), "finding repository")
}

// GetOwnedRepos returns all repos the user owns.
func (m *Model) GetOwnedRepos(u *User, search string) (RepositoryList, *errors.Error) {
	r := []*Repository{}
	where := "owner_id = ?"
	if search == "" {
		return RepositoryList(r), m.WrapError(m.Where(where, u.ID).Find(&r), "obtaining owned repositories")
	}
	return RepositoryList(r), m.WrapError(m.Where(where+" and name like ?", u.ID, "%"+search+"%").Find(&r), "obtaining owned repositories")

}

// GetVisibleReposForUser retrieves all repos the user can "see" in the
// database.
func (m *Model) GetVisibleReposForUser(u *User, search string) (RepositoryList, *errors.Error) {
	r, err := m.GetAllPublicRepos(search)
	if err != nil {
		return nil, err
	}

	r2, err := m.GetPrivateReposForUser(u, search)
	if err != nil {
		return nil, err
	}

	// reverse order to prefer private repos at the top
	return append(r2, r...), nil
}

// GetAllPublicRepos retrieves all repos that are not private
func (m *Model) GetAllPublicRepos(search string) (RepositoryList, *errors.Error) {
	// this call is probably a terrible idea for scaling things
	r := []*Repository{}
	where := "not private"
	if search == "" {
		return RepositoryList(r), m.WrapError(m.Where(where).Find(&r), "obtaining public repositories")
	}
	return RepositoryList(r), m.WrapError(m.Where(where+" and name like ?", "%"+search+"%").Find(&r), "obtaining public repositories")
}

// GetPrivateReposForUser retrieves all private repos that the user owns.
func (m *Model) GetPrivateReposForUser(u *User, search string) (RepositoryList, *errors.Error) {
	r := []*Repository{}

	where := "owner_id = ? and private"
	if search == "" {
		return RepositoryList(r), m.WrapError(
			m.Where(where, u.ID).
				Find(&r),
			"obtaining private repositories for user",
		)
	}
	return RepositoryList(r), m.WrapError(
		m.Where(where+" and name like ?", u.ID, "%"+search+"%").
			Find(&r),
		"obtaining private repositories for user",
	)
}

// GetRepositoryByName retrieves the repository by its unique name.
func (m *Model) GetRepositoryByName(name string) (*Repository, *errors.Error) {
	r := &Repository{}
	return r, m.WrapError(m.Where("name = ?", name).First(r), "obtain repository by name")
}

// AfterFind validates the output from the database before releasing it to the
// hook chain
func (r *Repository) AfterFind(tx *gorm.DB) error {
	if err := json.Unmarshal(r.GithubJSON, &r.Github); err != nil {
		return errors.New(err).Wrapf("reading github repository for id %d (%q)", r.ID, r.Name)
	}

	if err := r.Validate(false); err != nil {
		return errors.New(err).Wrapf("reading repository id %d (%q)", r.ID, r.Name)
	}

	return nil
}

// BeforeCreate just calls BeforeSave.
func (r *Repository) BeforeCreate(tx *gorm.DB) error {
	return r.BeforeSave(tx)
}

// BeforeSave is a gorm hook to marshal the token JSON before saving the record
func (r *Repository) BeforeSave(tx *gorm.DB) error {
	if err := r.Validate(true); err != nil {
		return errors.New(err).Wrapf("saving repository %q", r.Name)
	}

	var err error
	r.GithubJSON, err = json.Marshal(&r.Github)
	if err != nil {
		return errors.New(err).Wrapf("reading github repository for id %d (%q)", r.ID, r.Name)
	}

	return nil
}

// Validate validates the repository object
func (r *Repository) Validate(validOwner bool) *errors.Error {
	if r.Name == "" {
		return errors.New("name is empty")
	}

	if r.Github == nil {
		return errors.New("github content is nil")
	}

	if r.Name != r.Github.GetFullName() {
		return errors.New("github repository does not match repository name")
	}

	return nil
}

// Enabled is merely a predicate to determine if the repo can be used or not
func (r *Repository) Enabled() bool {
	return !r.Disabled && r.Owner != nil
}

// DisableRepository removes it from CI.
func (m *Model) DisableRepository(repo *Repository) *errors.Error {
	if !repo.Enabled() {
		return errors.New("repo is not enabled")
	}

	repo.Disabled = true
	return m.WrapError(m.Save(repo), "disabling repository")
}

// EnableRepository adds it to CI.
func (m *Model) EnableRepository(repo *Repository, owner *User) *errors.Error {
	if repo.Enabled() {
		return errors.New("repo is already enabled")
	}

	repo.Disabled = false
	repo.HookSecret = strings.TrimRight(base32.StdEncoding.EncodeToString(securecookie.GenerateRandomKey(24)), "=")
	repo.Owner = owner
	return m.WrapError(m.Save(repo), "enabling repository")
}

// AssignRepository assigns the repository to the user explicitly.
func (m *Model) AssignRepository(repo *Repository, owner *User) *errors.Error {
	repo.Owner = owner
	return m.WrapError(m.Save(repo), fmt.Sprintf("assigning repository to %q", owner.Username))
}
