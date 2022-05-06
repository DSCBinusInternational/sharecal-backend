package main

import (
	"sharecal-backend/controllers"
	"sharecal-backend/db"

	"github.com/gin-gonic/gin"
)

func main() {

	db.Init()
	r := gin.Default()
	r.GET("cal/:name", controllers.GetFunc)
	r.POST("cal/:name", controllers.PostFunc)
	r.PUT("cal/:name", controllers.PutFunc)

	r.Run()

}
