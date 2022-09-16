package routes

import (
	"assignment2/configs"
	"assignment2/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func LoadRoute() {
	db := configs.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.POST("/orders", inDB.CreateOrder)
	router.GET("/orders", inDB.GetOrders)
	router.PUT("/orders/:orderId", inDB.UpdateOrder)
	router.DELETE("/orders/:orderId", inDB.DeleteOrder)

	router.Run(fmt.Sprintf("localhost:%v", configs.Env.Port))
}
