package bootstrap

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func BootApplication() {
	mapUrls()
	router.Run(":8080")
}
