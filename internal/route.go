package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Route() {
	r := gin.Default()
	http.HandleFunc("/", HandleIndex)
	http.HandleFunc("/ws", HandleWebSocket)
	r.GET("/", func(c *gin.Context) {
		renderChat(c, chat)
	})

	r.POST("/", func(c *gin.Context) {
		addMessage(c, chat)
	})
}
