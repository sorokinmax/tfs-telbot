package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func tfsGitMerge(ctx *gin.Context) {

	jsonData, _ := io.ReadAll(ctx.Request.Body)
	log.Println(string(jsonData))

	/*var (
		gm  GitMerge
		msg string
	)
	err := ctx.BindJSON(&gm)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "json binding error"})
		return
	}

	msg += gm.Message.HTML + "\n\n"
	//msg += fmt.Sprintf(`<b>Client:</b> %s`, gm.Resource.Fields.BmClient)

	tgSendMessage(msg, cfg.Telegram.CIUpdatesChatID)
	*/
	ctx.JSON(http.StatusOK, gin.H{"status": "Done"})
}
