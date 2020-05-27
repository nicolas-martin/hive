package handler

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/nicolas-martin/hive/api/bot/client"
	"github.com/nicolas-martin/hive/api/config"
	"github.com/nicolas-martin/hive/api/model"
	"github.com/nicolas-martin/hive/api/repo"
)

// Handler ...
type Handler struct {
	repo        *repo.Repo
	slackClient *client.SlackClient
	frontendURL string
}

// NewHandler creates a new handler
func NewHandler(cfg *config.Config, repo *repo.Repo) *Handler {
	return &Handler{repo: repo, frontendURL: cfg.FrontEndURL}

}

// Ping responds with pong
func (h *Handler) Ping(c *gin.Context) {
	log.Println("inside the ping")
	c.JSON(200, gin.H{"message": "pong"})
}

// CreateUpdate creates an update request to be sent to users
func (h *Handler) CreateUpdate(c *gin.Context) {
	var update model.Update
	err := c.BindJSON(&update)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error binding json: %s", err.Error()))
		return
	}

	updateID, err := h.repo.AddUpate(&update)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error creating update: %s", err.Error()))
		return
	}

	userIDs := make([]string, 0)
	for _, v := range update.Usernames {
		userID, err := h.slackClient.GetUserID(v)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("error finding user %s", v))
			return
		}
		userIDs = append(userIDs, userID)
	}

	for _, v := range userIDs {
		ud := &model.UserUpdate{
			SlackUserID:  v,
			RecordingURL: "",
			UpdateID:     updateID,
		}

		udID, err := h.repo.AddUserUpate(ud)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("error creating user update %s", v))
			return
		}

		url := fmt.Sprintf("%s/record/%s", h.frontendURL, udID)
		msg := fmt.Sprintf("Please record a message for your update %s", url)
		h.slackClient.PostMessage(v, msg)
	}

	c.String(http.StatusOK, fmt.Sprintf("Created update successfully"))
	return
}

// Upload uploads a file
func (h *Handler) Upload(c *gin.Context) {
	user := c.PostForm("user")
	tenant := c.PostForm("tenant")

	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, "up/"+filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields user=%s and tenant=%s.", file.Filename, user, tenant))
	return
}
