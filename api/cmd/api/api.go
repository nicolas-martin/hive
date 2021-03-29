package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/nicolas-martin/hive/api/bot/client"
	"github.com/nicolas-martin/hive/api/config"
	"github.com/nicolas-martin/hive/api/handler"
	"github.com/nicolas-martin/hive/api/middleware"
	"github.com/nicolas-martin/hive/api/model"
	"github.com/nicolas-martin/hive/api/repo"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	logger := log.New()

	cfg, err := config.Load()
	if err != nil {
		log.Println(err)
		return
	}

	s := client.NewSlack(cfg)
	repo := repo.NewRepo(cfg)
	handler := handler.NewHandler(cfg, repo, s)

	gin.DefaultErrorWriter = logger.Writer()
	gin.DefaultWriter = logger.Writer()
	r := gin.Default()
	r.Use(middleware.Ginrus(logger, time.RFC3339, false))
	r.Use(cors.Default())
	r.GET("/ping", handler.Ping)
	r.POST("/upload", handler.Upload)
	r.POST("/update", handler.CreateUpdate)
	r.POST("/a", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"hello": "world"})
	})

	time.AfterFunc(100*time.Millisecond, func() {
		setup(repo, cfg, handler)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func setup(r *repo.Repo, cfg *config.Config, h *handler.Handler) {
	slackClient := client.NewSlack(cfg)
	users := []*model.User{{DisplayName: "martinni39", SlackUserID: ""}}

	for _, v := range users {
		userID, err := slackClient.GetUserID(v.DisplayName)
		if err != nil {
			log.Fatal(err)
		}
		v.SlackUserID = userID
		r.AddUser(v)
	}

	updateID, _ := r.AddUpate(&model.Update{Users: users})

	createUpdate := model.Update{UpdateID: updateID, Users: users}
	s, _ := json.Marshal(createUpdate)
	resp, err := http.Post("http://localhost:8080/update", "application/json", bytes.NewReader(s))
	if err != nil {
		log.Fatal("error sending post to update", err)
	}
	_ = resp

}
