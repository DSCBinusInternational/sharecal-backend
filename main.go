package main

import (
	"sharecal-backend/controllers"
	"sharecal-backend/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	db.Init()
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("cal/:name", controllers.GetFunc)
	r.GET("cal/:name/passcheck/:pass", controllers.PassCheckFunc)
	r.POST("cal/:name", controllers.PostFunc)
	r.Run()

}
