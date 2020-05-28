package client

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/nicolas-martin/hive/api/config"
	"github.com/slack-go/slack"
)

// SlackClient is a wrapper around slack-go/slack
type SlackClient struct {
	api *slack.Client
}

// NewSlack instantiate a new SlackClient
func NewSlack(cfg *config.Config) *SlackClient {
	api := slack.New(cfg.SlackToken)
	return &SlackClient{api: api}
}

// PostMessage sends a message to a given userID
func (s *SlackClient) PostMessage(userID string, message string) error {

	_, _, channelID, err := s.api.OpenIMChannel(userID)
	if err != nil {
		return err
	}

	_, _, err = s.api.PostMessage(channelID, slack.MsgOptionText(message, false))
	if err != nil {
		log.Fatal(err)
	}

	log.WithFields(log.Fields{
		"SENT":    channelID,
		"Message": message,
	}).Info()

	return err
}

// GetUserID queries the GetUsers API for the display name
func (s *SlackClient) GetUserID(displayName string) (string, error) {
	users, err := s.api.GetUsers()
	if err != nil {
		return "", err
	}

	for _, v := range users {
		if v.Profile.DisplayName == displayName {
			return v.ID, nil
		}

	}

	return "", fmt.Errorf("User not found")

}
