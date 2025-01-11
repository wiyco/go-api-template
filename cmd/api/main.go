package api

import (
	"github.com/gin-gonic/gin"
)

func InitApi() {
	router := gin.Default()

	ping(router)

	err := router.Run(":8000")
	if err != nil {
		panic(err)
	}
}
