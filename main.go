package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ngenohkevin/calorie_tracker/routes"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	router := gin.New()
	router.Use(gin.Logger())

	router.POST("/entry/create", routes.AddEntry)

	router.GET("/entries", routes.GetEntries)
	router.GET("/entry/:id/", routes.GetEntryById)
	router.GET("/ingredient/:ingredient", routes.GetEntriesByIngredient)

	router.PUT("/entry/update/:id", routes.UpdateEntry)
	router.PUT("/ingredient/update/:id", routes.UpdateIngredients)

	router.DELETE("/entry/delete/:id", routes.DeleteEntry)

	err := router.Run(":" + port)
	if err != nil {
		return
	}

}
