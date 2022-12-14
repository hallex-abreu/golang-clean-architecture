package http

import (
	"github.com/gin-gonic/gin"
	database "github.com/hallex-abreu/golang-clean-architecture/api/infra/database/gorm"
	customerControllers "github.com/hallex-abreu/golang-clean-architecture/api/infra/http/controllers/customer-controller"
	healthControllers "github.com/hallex-abreu/golang-clean-architecture/api/infra/http/controllers/health-controller"
)

func Init() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/health", healthControllers.Alive)
	router.GET("/customer", customerControllers.Index)
	router.POST("/customer", customerControllers.Store)

	database.Connection()

	router.Run()
}
