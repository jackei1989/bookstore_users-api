package bootstrap

import (
	"bookstoreUsersApi/logger"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router = gin.Default()
)

func BootApplication() {
	mapUrls()
	logger.Info("about to start application")
	if err := router.Run(":8080"); err != nil {
		log.Print(err)
	}
}
