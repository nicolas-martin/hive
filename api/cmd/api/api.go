package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicolas-martin/hive/api/config"
	"github.com/nicolas-martin/hive/api/handler"
	"github.com/nicolas-martin/hive/api/repo"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {

	cfg := config.Load()
	repo := repo.NewRepo(cfg)
	handler := handler.NewHandler(cfg, repo)

	r := gin.Default()
	r.Use(cors.Default())
	// r.Use(corsMiddleware())
	r.GET("/ping", handler.Ping)
	r.POST("/upload", handler.Upload)
	r.POST("/a", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"hello": "world"})
	})
	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
