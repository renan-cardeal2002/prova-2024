package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"prova-2024/handler"
)

func NewRouter(vehicleHandler *handler.VehicleHandler, accessoryHandler *handler.AccessoryHandler) *gin.Engine {
	route := gin.Default()
	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	router := route.Group("/api")
	vehicleRouter := router.Group("/vehicle")
	vehicleRouter.GET("", vehicleHandler.FindAll)
	vehicleRouter.GET("/:vehicleId", vehicleHandler.FindById)
	vehicleRouter.POST("", vehicleHandler.Create)
	vehicleRouter.PATCH("/:vehicleId", vehicleHandler.Update)
	vehicleRouter.DELETE("/:vehicleId", vehicleHandler.Delete)

	accessoryRouter := router.Group("/accessory")
	accessoryRouter.POST("", accessoryHandler.Create)
	accessoryRouter.DELETE("", accessoryHandler.Delete)

	return route
}
