package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tong-server-ja/controllers"
	"github.com/tong-server-ja/rabbitmq"
)

func init() {
	rabbitmq.RabbitMQInit()
}

func main() {
	router := gin.Default()
	router.POST("/create", controllers.Create)
	router.GET("/list", controllers.GetAll)
	router.GET("/get/:id", controllers.GetByID)
	router.GET("/delete/:id", controllers.Delete)
	router.GET("/undelete/:id", controllers.UnDelete)
	router.POST("/change.pass", controllers.ChangePassword)
	router.POST("/update", controllers.Updates)
	router.Run()
}
