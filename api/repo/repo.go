package repo

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/nicolas-martin/hive/api/config"
	"github.com/nicolas-martin/hive/api/model"
)

// Repo holds data in memory
type Repo struct {
	UserData       map[uuid.UUID]*model.User
	UserUpdateData map[uuid.UUID]*model.UserUpdate
	UpdateData     map[uuid.UUID]*model.Update
}

// NewRepo creater a new Repo
func NewRepo(cfg *config.Config) *Repo {
	userdata := make(map[uuid.UUID]*model.User)
	userupdatedata := make(map[uuid.UUID]*model.UserUpdate)
	updatedata := make(map[uuid.UUID]*model.Update)
	_ = cfg
	return &Repo{UserData: userdata, UserUpdateData: userupdatedata, UpdateData: updatedata}
}

// AddUser adds a user to be messaged
func (r *Repo) UpdateSlackUserIDByUserName(userName string, slackUserID string) error {
	for _, v := range r.UserData {
		if v.DisplayName == userName {
			v.SlackUserID = slackUserID
			return nil

		}

	}

	return fmt.Errorf("Cannot find user with userName %s", userName)

}

func (r *Repo) GetUserIDBySlackUserID(slackUserID string) (uuid.UUID, error) {
	var uuid uuid.UUID
	for k, v := range r.UserData {
		if v.SlackUserID == slackUserID {
			return k, nil
		}

	}
	return uuid, fmt.Errorf("Cannot get user with name %s", slackUserID)
}

// AddUser adds a user to be messaged
func (r *Repo) AddUser(user *model.User) (uuid.UUID, error) {
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
