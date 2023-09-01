package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// Helper method to set `db` to SQLite connection so we can make queries to `hackathons.db`
func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("hackathons.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Hackathon{}) // keeps our schema up to date
	if err != nil {
		return
	}

	db = database
}

// Defines what a `Hackathon` is
type Hackathon struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Date     string `json:"date"`
	Url      string `json:"url"`
	Location string `json:"location"`
}

// GET /hackathons
// Get all hackathons
func getHackathons(c *gin.Context) {
	// Code goes here
}

// GET /hackathons/:id
// Get hackathon by ID
func getHackathonById(c *gin.Context) {
	// Code goes here
}

// POST /hackathons
// Create a hackathon
func createHackathon(c *gin.Context) {
	// Code goes here
}

// PATCH /hackathons/:id
// Update a hackathon
func updateHackathon(c *gin.Context) {
	// Code goes here
}

// DELETE /hackathons/:id
// Delete a hackathon
func deleteHackathon(c *gin.Context) {
	// Code goes here
}

func main() {
	router := gin.Default()

	ConnectDatabase()

	// router.GET("/hackathons", getHackathons)
	// router.GET("/hackathons/:id", getHackathonById)
	// router.POST("/hackathons", createHackathon)
	// router.PATCH("/hackathons/:id", updateHackathon)
	// router.DELETE("/hackathons/:id", deleteHackathon)

	router.Run("localhost:8080")
}
