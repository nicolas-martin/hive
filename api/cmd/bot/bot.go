package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/nicolas-martin/hive/api/bot/client"
	"github.com/nicolas-martin/hive/api/config"
	"github.com/nicolas-martin/hive/api/model"
	"github.com/nicolas-martin/hive/api/repo"
)

func main() {
	cfg := config.Load()
	s := client.NewSlack(cfg)
	r := repo.NewRepo(cfg)
	r.AddUser(&model.User{DisplayName: "martinni39", SlackUserID: ""})

	for _, v := range r.UserData {
		userID, err := s.GetUserID(v.DisplayName)
		if err != nil {
			log.Fatal(err)
		}

		err = s.PostMessage(userID, "Hello world")
		if err != nil {
			log.Fatal(err)
		}
	}

}
