package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"url-shortener/internal/store"
	"url-shortener/internal/utils"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("internal/templates/*")

	memMap := make(map[string]string)
	memory := store.StoreBackEnd{&memMap}

	router.POST("/shorten", func(c *gin.Context) {
		url := c.PostForm("url")

		alphaNum := utils.RandSeq(3)

		for ok := true; ok; ok = memory.IsExist(alphaNum) {
			alphaNum = utils.RandSeq(3)
		}
		memory.Add(alphaNum, url)

		c.JSON(http.StatusOK, gin.H{
			"shortened_url": fmt.Sprintf("http://localhost:8080/%v", alphaNum),
		})
	})

	router.GET("*url", func(c *gin.Context) {

		shortened_url := strings.Replace(c.Param("url"), "/", "", 1)
		if memory.IsExist(shortened_url) {
			c.Redirect(http.StatusFound, memory.Get(shortened_url))
		} else {
			c.HTML(http.StatusNotFound, "404.tmpl", gin.H{
				"message": "404 not found",
			})
		}
	})

	router.Run(":8080")
}
