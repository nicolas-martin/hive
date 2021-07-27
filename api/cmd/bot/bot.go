package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"

	"github.com/nicolas-martin/hive/api/bot/client"
	"github.com/nicolas-martin/hive/api/config"
	"github.com/nicolas-martin/hive/api/model"
	"github.com/nicolas-martin/hive/api/repo"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	s := client.NewSlack(cfg)
	r := repo.NewRepo(cfg)
	r.AddUser(&model.User{DisplayName: "martinni39", SlackUserID: ""})

	userID, err := s.GetUserID("martinni39")
	if err != nil {
		log.Fatal(err)
	}

	// TODO: try out more complex messages
	// err = s.PostMessage(userID, "Hello world")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	btest := &slack.Blocks{}
	err = btest.UnmarshalJSON(blocksStr)
	if err != nil {
		log.Fatal(err)
	}

	err = s.PostBlock(userID, btest)
	if err != nil {
		log.Fatal(err)
	}

}

var blocksStr = []byte(`{"blocks":[{"type":"section","text":{"type":"mrkdwn","text":"This is a section block with a button."},"accessory":{"type":"button","text":{"type":"plain_text","text":"Click Me"},"value":"click_me_123","action_id":"button"}},{"type":"actions","block_id":"actionblock789","elements":[{"type":"button","text":{"type":"plain_text","text":"Primary Button"},"style":"primary","value":"click_me_456"},{"type":"button","text":{"type":"plain_text","text":"Link Button"},"url":"https://api.slack.com/block-kit"}]}]}`)
