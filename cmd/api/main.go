package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		fmt.Println(c.Request.RemoteAddr)
		fmt.Println(c.Request.Host)
		fmt.Println(c.Request.URL)
		c.SetCookie("fahim", "fahim", 3600000000, "/", "localhost", false, false)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Server is running on port 5000",
			"data":    nil,
		})

	})

	addr := fmt.Sprint(":5000")
	r.Run(addr)

}
