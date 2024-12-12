package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/valeennmendez/to-do/connection"
	"github.com/valeennmendez/to-do/models"
	"github.com/valeennmendez/to-do/routes"
)

func main() {

	connection.Connection()

	connection.DB.AutoMigrate(&models.Task{})
	connection.DB.AutoMigrate(&models.User{})

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:3000"},
		AllowMethods:     []string{"GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authoritazion"},
		AllowCredentials: true,
	}))

	r.POST("/register", routes.RegisterUser)
	r.POST("/login", routes.Login)
	r.POST("/create-task", routes.CreateTask)
	r.DELETE("/delete-task/:id", routes.DeleteTask)
	r.GET("/", routes.GetAllTask)
	r.GET("/task", routes.GetTask)

	r.Run(":8080")

}
