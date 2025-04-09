package main 


import (
	
  
	"github.com/gin-gonic/gin"
  )


func main() {
	server := gin.Default()

	server.GET("/events", getEvents ) //* get,post,put,delete,patch

	server.Run(":8080") //* localhost:8080

}


func getEvents(c *gin.Context) {
	// Logic to retrieve events from the database or any data source
	events := []string{"Event 1", "Event 2", "Event 3"}
	c.JSON(200, events)
}