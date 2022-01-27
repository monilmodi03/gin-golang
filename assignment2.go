package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.html", nil)
	})

	r.POST("/", func(c *gin.Context) {

		namevalue := c.PostForm("name")
		phonevalue := c.PostForm("phone")
		gendervalue := c.PostForm("gender")
		addressvalue := c.PostForm("address")
		emailvalue := c.PostForm("email")
		datevalue := c.PostForm("birthday")

		c.JSON(200, gin.H{
			"Name":    namevalue,
			"Phone":   phonevalue,
			"Gender":  gendervalue,
			"Address": addressvalue,
			"Email":   emailvalue,
			"DOB":     datevalue})
	})

	r.Run(":3000")
}
