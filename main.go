package main

import (
	"net/http"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)


func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents ) //* get,post,put,delete,patch
	server.POST("/events", createEvent)

	server.Run(":8080") //* localhost:8080

}


func getEvents(c *gin.Context) {
	// Logic to retrieve events from the database or any data source
	events := models.GetAllEvents() 
	
	c.JSON(200, events)
}


func createEvent(c *gin.Context){
	var event models.Event
	err := c.ShouldBindJSON(&event)


	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	event.ID = 1 
	event.UserID = 1

	event.Save()

	c.JSON(http.StatusCreated,gin.H{"message": "Event created successfully", "event": event})


}