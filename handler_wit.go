package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func tfsCRCreated(ctx *gin.Context) {
	var (
		cr  WITCRCreate
		msg string
	)
	err := ctx.BindJSON(&cr)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "json binding error"})
		return
	}

	msg += cr.Message.HTML + "\n\n"
	msg += fmt.Sprintf(`<b>Client:</b> %s`, cr.Resource.Fields.BmClient)

	tgSendMessage(msg, cfg.Telegram.CRCreatesChatID)

	ctx.JSON(http.StatusOK, gin.H{"status": "Done"})
}

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
	msg += fmt.Sprintf(`<b>Priority:</b> %d`, ci.Resource.Fields.MicrosoftVSTSCommonPriority) + "\n"
	msg += fmt.Sprintf(`<b>Client:</b> %s`, ci.Resource.Fields.BmClient)

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
	// 200 because TFS fails the webhook at 4xx codes and disables it
	if !ci.Resource.Revision.Fields.BmZendesk {
		log.Println("Decline: BmZendesk is false")
		ctx.JSON(http.StatusOK, gin.H{"status": "Decline bacause BmZendesk is false"})
		return
	}

	msg += ci.Message.HTML + "\n\n"
	msg += `<b>UPDATED</b>`

	tgSendMessage(msg, cfg.Telegram.CIUpdatesChatID)

	ctx.JSON(http.StatusOK, gin.H{"status": "Done"})
}
