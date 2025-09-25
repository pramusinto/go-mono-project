package main

import (
	"context"
	"net/http"
	"producer-service/internal/config"
	"producer-service/internal/kafka"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	producer := kafka.NewProducer(cfg.KafkaBroker, cfg.KafkaTopic)
	r := gin.Default()

	r.POST("/publish", func(c *gin.Context) {
		var req struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := producer.Publish(context.Background(), []byte(req.Key), []byte(req.Value))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "publish"})
	})
	r.Run(":8800")
}
