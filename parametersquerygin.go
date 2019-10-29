package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	router := gin.Default()

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		fname := c.DefaultQuery("firstname", "Guest")
		lname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", fname, lname)
	})
	router.Run(":8080")
}
