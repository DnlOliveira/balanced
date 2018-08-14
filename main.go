package main

import (
	"net/http"

	"github.com/DnlOliveira/balanced/pkg"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var m *pkg.Mongo

func main() {
	m = &pkg.Mongo{"localhost:27017", "balanced"}
	m.Connect()

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./build", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		api.GET("/user/:email", getUser) //retrieve user and budgets
		api.POST("/user", addUser)       //add user to db
		// api.POST("/budget", updateBudget) //add or update budget
		// api.POST("/wallet", updateWallet) //add or update wallet
	}

	// Start and run the server
	router.Run(":3000")
}

func getUser(c *gin.Context) {
	key := "email"
	p := c.Param("email")
	user := m.Find(key, &p)

	c.JSON(http.StatusOK, gin.H{
		"Status": "Success",
		"user":   user,
	})
}

func addUser(c *gin.Context) {
	email := c.Param("email")
	user := m.Find("email", &email)

	c.JSON(http.StatusOK, gin.H{
		"Status": "Success",
		"user":   user,
	})
}
