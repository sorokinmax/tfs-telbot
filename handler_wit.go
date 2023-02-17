package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func tfsCICreated(ctx *gin.Context) {
	var (
		ci  WITCICreate
		msg string
	)
	err := ctx.BindJSON(&ci)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "json binding error"})
		return
	}

	msg += ci.Message.HTML + "\n\n"
	msg += fmt.Sprintf(`<b>Priority:</b> %d`, ci.Resource.Fields.MicrosoftVSTSCommonPriority)

	tgSendMessage(msg, cfg.Telegram.CICreatesChatID)

	ctx.JSON(http.StatusOK, gin.H{"status": "Done"})
}

func tfsCIUpdated(ctx *gin.Context) {
	var (
		ci  WITCIUpdate
		msg string
	)
	err := ctx.BindJSON(&ci)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "json binding error"})
		return
	}
	if !ci.Resource.Revision.Fields.BmZendesk {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"cause": "BmZendesk is false"})
		return
	}

	msg += ci.Message.HTML + "\n\n"
	msg += `<b>UPDATED</b>` + "\n\n"
	msg += `Ну че пидрилы слепые`

	tgSendMessage(msg, cfg.Telegram.CIUpdatesChatID)

	ctx.JSON(http.StatusOK, gin.H{"status": "Done"})
}
