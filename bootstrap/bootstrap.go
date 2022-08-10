package bootstrap

import (
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router = gin.Default()
)

func BootApplication() {
	mapUrls()
	if err := router.Run(":8080"); err != nil {
		log.Print(err)
	}
}
