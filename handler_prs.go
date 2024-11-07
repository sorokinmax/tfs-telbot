package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
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

	// Замена TFS-ных тэгов пользователей в описании (@<EBF4FBB9-6B01-4178-B118-6CEB7BD5118C>)
	regex := regexp.MustCompile("@<.*>")
	Str := "Jesus Christ"
	description := regex.ReplaceAllString(pr.Resource.Description, Str)

	msg = "🎀🎀🎀\n\n"
	msg += fmt.Sprintf(`<b>Title:</b> <a href='%s'>%s</a>`, pr.Resource.Links.Web.Href, pr.Resource.Title) + "\n"
	msg += fmt.Sprintf(`<b>Author:</b> %s`, pr.Resource.CreatedBy.DisplayName) + "\n"
	msg += fmt.Sprintf(`<b>Reviewers:</b> %s`, strings.Join(reviewers[:], ",")) + "\n"
	msg += fmt.Sprintf(`<b>Source:</b> [%s] %s`, pr.Resource.Repository.Name, strings.TrimPrefix(pr.Resource.SourceRefName, "refs/heads/")) + "\n"
	msg += fmt.Sprintf(`<b>Target:</b> [%s] %s`, pr.Resource.Repository.Name, strings.TrimPrefix(pr.Resource.TargetRefName, "refs/heads/")) + "\n"
	msg += fmt.Sprintf(`<b>Description:</b> %s`, description)

	switch pr.Resource.Repository.Name {
	case cfg.Tfs.BackendRepo:
		tgSendMessage(msg, cfg.Telegram.PRBackendCreatesChatID)
	case cfg.Tfs.FrontendRepo:
		tgSendMessage(msg, cfg.Telegram.PRFrontendCreatesChatID)
	case cfg.Tfs.AndroidRepo:
		tgSendMessage(msg, cfg.Telegram.PRAndroidCreatesChatID)
	case cfg.Tfs.IosRepo:
		tgSendMessage(msg, cfg.Telegram.PRIosCreatesChatID)
	default:
		tgSendMessage(msg, cfg.Telegram.GeneralDevChatID)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Done"})
}
