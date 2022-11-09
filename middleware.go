package main

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

func ginBodyLogMiddleware(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	log.Printf("Request dump: %s\n", string(body))

	c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
}
