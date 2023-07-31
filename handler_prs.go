package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func tfsPRCreated(ctx *gin.Context) {
	var (
		pr  PRCreate
		msg string
	)
	err := ctx.BindJSON(&pr)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "json binding error"})
		return
	}

	var reviewers []string
	for _, reviewer := range pr.Resource.Reviewers {
		reviewers = append(reviewers, reviewer.DisplayName)
	}

	msg = "🎀🎀🎀\n\n"
	msg += fmt.Sprintf(`<b>Title:</b> <a href='%s'>%s</a>`, pr.Resource.Links.Web.Href, pr.Resource.Title) + "\n"
	msg += fmt.Sprintf(`<b>Author:</b> %s`, pr.Resource.CreatedBy.DisplayName) + "\n"
	msg += fmt.Sprintf(`<b>Reviewers:</b> %s`, strings.Join(reviewers[:], ",")) + "\n"
	msg += fmt.Sprintf(`<b>Source:</b> %s`, pr.Resource.SourceRefName) + "\n"
	msg += fmt.Sprintf(`<b>Target:</b> %s`, pr.Resource.TargetRefName) + "\n"
	msg += fmt.Sprintf(`<b>Description:</b> %s`, pr.Resource.Description)

	tgSendMessage(msg, cfg.Telegram.PRCreatesChatID)

	ctx.JSON(http.StatusOK, gin.H{"status": "Done"})
}
