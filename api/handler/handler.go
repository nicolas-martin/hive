package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

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
func NewHandler(cfg *config.Config, repo *repo.Repo, sl *client.SlackClient) *Handler {
	return &Handler{repo: repo, frontendURL: cfg.FrontEndURL, slackClient: sl}

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

	for _, v := range update.Users {
		slackUserID, err := h.slackClient.GetUserID(v.DisplayName)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("error finding user %s", v))
			return
		}

		url := fmt.Sprintf("%s/record/%s/%s", h.frontendURL, v.UserID, updateID)
		msg := fmt.Sprintf("Please record a message for your update %s", url)
		h.slackClient.PostMessage(slackUserID, msg)
	}

	c.String(http.StatusOK, "Created update successfully")
}

// Upload uploads a file
func (h *Handler) Upload(c *gin.Context) {
	updateIDstr := c.PostForm("updateid")
	updateID, err := uuid.Parse(updateIDstr)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Invalid Update %s", updateIDstr))
		log.WithFields(log.Fields{
			"Err": errors.Wrap(err, fmt.Sprintf("invalid UpdateID %s", updateIDstr)),
		}).Error()
		return
	}

	userIDstr := c.PostForm("userid")
	userID, err := uuid.Parse(userIDstr)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Invalid UserID %s", userIDstr))
		log.WithFields(log.Fields{
			"Err": errors.Wrap(err, fmt.Sprintf("invalid UserID %s", userIDstr)),
		}).Error()
		return
	}

	// Check if the updateID received is a valid update
	update, err := h.repo.GetUpdateAndVerifyUser(updateID, userID)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		log.WithFields(log.Fields{
			"Err": err,
		}).Error()
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		log.WithFields(log.Fields{
			"Err": errors.Wrap(err, "get form err"),
		}).Error()
		return
	}

	// filename := filepath.Base(file.Filename)
	path := fmt.Sprintf("/Users/nmartin/go/src/github.com/nicolas-martin/hive/up/%s", update.UpdateID)
	fullPath := fmt.Sprintf("%s/%s.webm", path, updateID)
	err = os.MkdirAll(path, os.ModePerm)

	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error creating the upload folder: %s", err.Error()))
		log.WithFields(log.Fields{
			"Err": errors.Wrap(err, "Error creating the upload folder"),
		}).Error()
		return
	}

	log.WithFields(log.Fields{
		"path": fullPath,
	}).Info("Created Folder")

	if err := c.SaveUploadedFile(file, fullPath); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		log.WithFields(log.Fields{
			"Err": errors.Wrap(err, "upload file error"),
		}).Error()
		return
	}

	log.WithFields(log.Fields{
		"updateID": updateID,
	}).Info("UploadedFile")

	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with updateID=%s.", file.Filename, updateID))

	ud := &model.UserUpdate{
		UserID:       userID,
		RecordingURL: fullPath,
		UpdateID:     updateID,
	}

	udID, err := h.repo.AddUserUpate(ud)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error creating user update %s", udID))
		return
	}

	log.WithFields(log.Fields{
		"UserUpdateID": udID,
		"UserID":       userID,
		"UpdateID":     updateID,
	}).Info("Saved user update.")

	// Check if all the members have completed their update
	if completed, err := h.repo.CheckForCompletedUpdate(update.UpdateID); completed {
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		users, err := h.repo.GetUsersByUpdateID(update.UpdateID)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		for _, v := range users {
			err := h.slackClient.PostMessage(v.SlackUserID, fmt.Sprintf("Your team's update is completed at %s/%s", h.frontendURL, update.UpdateID))
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}

		}

	}
}
