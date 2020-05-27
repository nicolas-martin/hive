package repo

import (
	"github.com/google/uuid"
	"github.com/nicolas-martin/hive/api/config"
	"github.com/nicolas-martin/hive/api/model"
)

// Repo holds data in memory
type Repo struct {
	UserData       map[uuid.UUID]string
	UserUpdateData map[uuid.UUID]*model.UserUpdate
	UpdateData     map[uuid.UUID]*model.Update
}

// NewRepo creater a new Repo
func NewRepo(cfg *config.Config) *Repo {
	d := make(map[uuid.UUID]string)
	_ = cfg
	return &Repo{UserData: d}
}

// AddUser adds a user to be messaged
func (r *Repo) AddUser(user string) (uuid.UUID, error) {
	id := uuid.New()
	r.UserData[id] = user
	return id, nil
}

// AddUpate adds a user to be messaged
func (r *Repo) AddUpate(update *model.Update) (uuid.UUID, error) {
	id := uuid.New()
	r.UpdateData[id] = update
	return id, nil
}

func (r *Repo) AddUserUpate(userUpdate *model.UserUpdate) (uuid.UUID, error) {
	id := uuid.New()
	r.UserUpdateData[id] = userUpdate
	return id, nil
}
