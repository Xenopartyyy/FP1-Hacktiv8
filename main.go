package main

import (
	todoscontroller "FP1-Hacktiv8/controllers/todoscontroller"
	_ "FP1-Hacktiv8/docs"
	db "FP1-Hacktiv8/infra/db"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()
	db.ConnDB()

	//Read all todos
	r.GET("/todos", todoscontroller.Index)

	//Read todos by ID
	r.GET("/todos/:id", todoscontroller.Show)

	//Create new todos
	r.POST("/todos", todoscontroller.Create)

	//Update todos
	r.PUT("/todos/:id", todoscontroller.Update)

	//Delete todos by id
	r.DELETE("/todos/:id", todoscontroller.Delete)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":8080")
}
