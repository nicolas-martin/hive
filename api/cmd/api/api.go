package main

import (
	"fmt"
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

	repo := repo.NewRepo(cfg)
	handler := handler.NewHandler(cfg, repo)
	log.Debug("????A?S?SAFSDF")

	gin.DefaultErrorWriter = logger.Writer()
	gin.DefaultWriter = logger.Writer()
	r := gin.Default()
	r.Use(middleware.Ginrus(logger, time.RFC3339, false))
	r.Use(cors.Default())
	r.GET("/ping", handler.Ping)
	r.POST("/upload", handler.Upload)
	r.POST("/a", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"hello": "world"})
	})

	setup(repo, cfg, handler)

	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func setup(r *repo.Repo, cfg *config.Config, h *handler.Handler) {
	slackClient := client.NewSlack(cfg)
	users := []*model.User{{DisplayName: "martinni39", SlackUserID: ""}}

	userNames := []string{}

	for _, v := range users {
		userNames = append(userNames, v.DisplayName)
		userID, err := slackClient.GetUserID(v.DisplayName)
		if err != nil {
			log.Fatal(err)
		}
		v.SlackUserID = userID
		r.AddUser(v)
	}

	updateID, _ := r.AddUpate(&model.Update{Users: users})

	for _, user := range users {
		userID, _ := r.GetUserIDBySlackUserID(user.SlackUserID)
		ud := &model.UserUpdate{
			UserID:       userID,
			RecordingURL: "",
			UpdateID:     updateID,
		}

		udID, _ := r.AddUserUpate(ud)
		url := fmt.Sprintf("%s/record/%s", cfg.FrontEndURL, udID)
		msg := fmt.Sprintf("Please record a message for your update %s", url)
		slackClient.PostMessage(user.SlackUserID, msg)
	}

}
