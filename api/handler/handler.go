package handler

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

		// NOTE: Update our record of the user
		// h.repo.UpdateSlackUserIDByUserName(v, userID)
		userIDs = append(userIDs, userID)
	}

	for _, v := range userIDs {
		// TODO: this will probably be done somewhere else
		userID, _ := h.repo.GetUserIDBySlackUserID(v)

		ud := &model.UserUpdate{
			UserID:       userID,
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
	updateIDstr := c.PostForm("id")
	updateID, err := uuid.Parse(updateIDstr)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Invalid UserUpdate %s", updateIDstr))
		return
	}

	// NOTE: Check if the ID received is a valid update
	v, ok := h.repo.UserUpdateData[updateID]
	if !ok {
		c.String(http.StatusBadRequest, fmt.Sprintf("UserUpdateID not found"))
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	filename := filepath.Base(file.Filename)
	path := fmt.Sprintf("up/%s", filename)
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	// NOTE: Update recording URL
	v.RecordingURL = path

	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with updateID=%s.", file.Filename, updateID))
	return
}
