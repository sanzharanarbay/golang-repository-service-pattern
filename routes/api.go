package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzharanarbay/repository-service-pattern/controllers"
	database "github.com/sanzharanarbay/repository-service-pattern/db"
	"github.com/sanzharanarbay/repository-service-pattern/repositories"
	"github.com/sanzharanarbay/repository-service-pattern/services"
)

func ApiRoutes(prefix string, router *gin.Engine) {
	db := database.ConnectDB()
	apiGroup := router.Group(prefix)
	{
		dashboard := apiGroup.Group("/dashboard/items")
		{
			itemRepo := repositories.NewItemRepository(db)
			itemService := services.NewItemService(itemRepo)
			itemController := controllers.NewItemController(itemService)

			dashboard.GET("/all", itemController.GetItemList)
			dashboard.GET("/:id", itemController.GetItem)
			dashboard.POST("/create", itemController.CreateItem)
			dashboard.PUT("/update/:id", itemController.UpdateItem)
			dashboard.DELETE("/delete/:id", itemController.DeleteItem)
		}
	}
}
