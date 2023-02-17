package main

import (
	"bytes"
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

func ginBodyLogMiddleware(c *gin.Context) {
	if cfg.Web.Debug {
		body, _ := io.ReadAll(c.Request.Body)
		log.Printf("Request dump: %s\n", string(body))

		c.Request.Body = io.NopCloser(bytes.NewReader(body))
	}
}
