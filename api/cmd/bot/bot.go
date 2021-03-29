package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/nicolas-martin/hive/api/bot/client"
	"github.com/nicolas-martin/hive/api/config"
	"github.com/nicolas-martin/hive/api/model"
	"github.com/nicolas-martin/hive/api/repo"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Println(err)
		return
	}

	s := client.NewSlack(cfg)
	r := repo.NewRepo(cfg)
	r.AddUser(&model.User{DisplayName: "martinni39", SlackUserID: ""})

	userID, err := s.GetUserID("martinni39")
	if err != nil {
		log.Fatal(err)
	}

	// TODO: try out more complex messages
	err = s.PostMessage(userID, "Hello world")
	if err != nil {
		log.Fatal(err)
	}

}
