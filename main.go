package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"userAuth/src/handlers"
)

func main() {
	r := gin.New()


	r.LoadHTMLGlob("src/html/*")

	r.GET("/", handlers.LoginGet)
	r.GET("/login", handlers.LoginGet)

	r.GET("/logout", handlers.Logout)
	r.POST("/", handlers.LoginPost)


	log.Fatal(r.Run(":80"))
}
