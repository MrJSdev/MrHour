package main

import (
	"MrHour/http"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, Brother")
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/hours", http.GetHours)
	router.GET("/portfolio", http.GetPortfolios)
	router.GET("/messages", http.GetMessages)
	router.POST("/power", http.GetPower)
	router.POST("/hour/save", http.SaveHour)
	router.POST("/message/send", http.SaveMessage)
	router.Run()
}
