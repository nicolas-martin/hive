package repo

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/nicolas-martin/hive/api/config"
	"github.com/nicolas-martin/hive/api/model"
)

// Repo holds data in memory
type Repo struct {
	userData       []*model.User
	userupdateData []*model.UserUpdate
	updateData     []*model.Update
}

// NewRepo creater a new Repo
func NewRepo(cfg *config.Config) *Repo {
	userData := make([]*model.User, 0)
	userupdateData := make([]*model.UserUpdate, 0)
	updateData := make([]*model.Update, 0)
	_ = cfg
	return &Repo{userData: userData, userupdateData: userupdateData, updateData: updateData}
}

// AddUser adds a user to be messaged
func (r *Repo) UpdateSlackUserIDByUserName(userName string, slackUserID string) error {
	for _, v := range r.userData {
		if v.DisplayName == userName {
			v.SlackUserID = slackUserID
			return nil
		}
	}

	return fmt.Errorf("cannot find user with userName %s", userName)

}

func (r *Repo) GetUpdate(updateID uuid.UUID) (*model.Update, error) {
	for _, v := range r.updateData {
		if v.UpdateID == updateID {
			return v, nil
		}

	}
	return nil, fmt.Errorf("cannot get update with ID %s", updateID)
}

func (r *Repo) GetUpdateAndVerifyUser(updateID, userID uuid.UUID) (*model.Update, error) {

	// make sure user hasn't posted an update already
	for _, v := range r.userupdateData {
		if v.UserID == userID && v.UpdateID == updateID {
			return nil, errors.New("user already submitted user update")
		}
	}

	for _, v := range r.updateData {
		if v.UpdateID == updateID {
			// make sure the user is in the update user list
			for _, u := range v.Users {
				if u.UserID == userID {
					return v, nil
				}
			}
			return nil, errors.New("user not authorized to make update")
		}

	}
	return nil, fmt.Errorf("cannot get update with ID %s", updateID)
}

func (r *Repo) GetUserUpdate(userUpdateID uuid.UUID) (*model.UserUpdate, error) {
	for _, v := range r.userupdateData {
		if v.UserUpdateID == userUpdateID {
			return v, nil
		}

	}
	return nil, fmt.Errorf("cannot get userUpdate with ID %s", userUpdateID)
}

func (r *Repo) GetUserIDBySlackUserID(slackUserID string) (uuid.UUID, error) {
	var uuid uuid.UUID
	for _, v := range r.userData {
		if v.SlackUserID == slackUserID {
			return v.UserID, nil
		}

	}
	return uuid, fmt.Errorf("cannot get user with name %s", slackUserID)
}

// AddUser adds a user to be messaged
func (r *Repo) AddUser(user *model.User) (uuid.UUID, error) {
	id := uuid.New()

	user.UserID = id
	r.userData = append(r.userData, user)
	return id, nil
}

// AddUpate adds a user to be messaged
func (r *Repo) AddUpate(update *model.Update) (uuid.UUID, error) {
	id := uuid.New()
	update.UpdateID = id
	r.updateData = append(r.updateData, update)
	return id, nil
}

func (r *Repo) AddUserUpate(userUpdate *model.UserUpdate) (uuid.UUID, error) {
	id := uuid.New()
	userUpdate.UserUpdateID = id
	r.userupdateData = append(r.userupdateData, userUpdate)
	return id, nil
}

// GetUsersByUpdateID returns users from updateID. MAYBE I DON'T NEED?
func (r *Repo) GetUsersByUpdateID(updateID uuid.UUID) ([]*model.User, error) {
	for _, v := range r.updateData {
		if v.UpdateID == updateID {
			return v.Users, nil
		}

	}
	return nil, fmt.Errorf("cannot get Users for update with ID %s", updateID)

}

func (r *Repo) GetUserUpdateByUserIDandUpdateID(updateID uuid.UUID, userID uuid.UUID) (*model.UserUpdate, error) {
	for _, v := range r.userupdateData {
		if v.UserID == userID && v.UpdateID == updateID {
			return v, nil

		}

	}
	return nil, fmt.Errorf("cannot get UserUpdate for updateID %s and userID %s", updateID, userID)

}

func (r *Repo) CheckForCompletedUpdate(updateID uuid.UUID) (bool, error) {
	update, err := r.GetUpdate(updateID)
	if err != nil {
		return false, err
	}

	for _, v := range update.Users {
		user, err := r.GetUserUpdateByUserIDandUpdateID(updateID, v.UserID)
		if err != nil {
			return false, err
		}

		if len(user.RecordingURL) == 0 {
			return false, nil

		}

	}

	return true, nil
}
