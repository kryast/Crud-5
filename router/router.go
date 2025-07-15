package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kryast/Crud-5.git/handlers"
)

func SetupRouter(orderItemHandler *handlers.OrderItemHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/order-items", orderItemHandler.Create)
	r.GET("/order-items", orderItemHandler.GetAll)
	r.GET("/order-items/:id", orderItemHandler.GetByID)

	return r
}
