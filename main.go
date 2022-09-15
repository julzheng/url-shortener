package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
	"url-shortener/internal/store"
	. "url-shortener/internal/types"
	"url-shortener/internal/utils"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("internal/templates/*")

	memMap := make(map[string]URLInfo)
	memory := store.StoreBackEnd{&memMap}

	router.POST("/shorten", func(c *gin.Context) {
		url := c.PostForm("url")

		alphaNum := utils.RandSeq(3)

		for ok := true; ok; ok = memory.IsExist(alphaNum) {
			alphaNum = utils.RandSeq(3)
		}
		memory.Add(alphaNum, URLInfo{url, 0, time.Now()})

		c.JSON(http.StatusOK, gin.H{
			"shortened_url": fmt.Sprintf("http://localhost:8080/%v", alphaNum),
		})
	})

	router.GET(":url", func(c *gin.Context) {

		shortened_url := strings.Replace(c.Param("url"), "/", "", 1)
		if memory.IsExist(shortened_url) {
			urlInfo := memory.Get(shortened_url)
			urlInfo.RedirectCount += 1
			memory.Save(shortened_url, urlInfo)
			fmt.Printf("Redirect count: %v times, Created at: %v", urlInfo.RedirectCount, urlInfo.CreatedAt)

			c.Redirect(http.StatusFound, urlInfo.URL)
		} else {
			c.HTML(http.StatusNotFound, "404.tmpl", gin.H{
				"message": "404 not found",
			})
		}
	})

	router.GET(":url/stats", func(c *gin.Context) {
		shortened_url := strings.Replace(c.Param("url"), "/", "", 1)
		if memory.IsExist(shortened_url) {
			urlInfo := memory.Get(shortened_url)

			c.JSON(http.StatusOK, gin.H{
				"created_at":     urlInfo.CreatedAt,
				"redirect_count": urlInfo.RedirectCount,
			})
		}
	})

	return router
}

func main() {
	router := SetupRouter()
	router.Run(":8080")
}
