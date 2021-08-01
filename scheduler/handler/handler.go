package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicolas-martin/hive/scheduler/config"
	"github.com/nicolas-martin/hive/scheduler/model"
	"github.com/nicolas-martin/hive/scheduler/repo"
	"github.com/nicolas-martin/hive/scheduler/worker"
)

// Handler ...
type Handler struct {
	repo   *repo.Repo
	worker *worker.Worker
}

// NewHandler creates a new handler
func NewHandler(cfg *config.Config, repo *repo.Repo, w *worker.Worker) *Handler {
	return &Handler{repo: repo, worker: w}

}

// Add adds a schedule
func (h *Handler) Add(c *gin.Context) {
	var schedule model.Schedule

	err := c.BindJSON(&schedule)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error binding json: %s", err.Error()))
		return
	}
	c.JSON(200, gin.H{"message": "pong"})
}
