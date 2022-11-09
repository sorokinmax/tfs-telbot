package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

func tfsBuildReport(ctx *gin.Context) {
	var (
		build Build
		msg   string
	)
	err := ctx.BindJSON(&build)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "json binding error"})
		return
	}

	if slices.Contains(cfg.ExcludedPipelines.Build, build.Resource.Definition.Name) {
		ctx.AbortWithStatusJSON(
			http.StatusPreconditionFailed,
			gin.H{"Response": fmt.Sprintf(`Pipeline "%s" is exluded`, build.Resource.Definition.Name)},
		)
		return
	}

	switch build.Resource.Result {
	case "succeeded":
		msg += "✅✅✅\n🏗🏗🏗\n\n"
	case "failed":
		msg += "❌❌❌\n🏗🏗🏗\n\n"
	case "partiallySucceeded":
		msg += "⚠️⚠️⚠️\n🏗🏗🏗\n\n"
	}

	msg += fmt.Sprintf(`<b>Definition:</b> <a href='%s%s/_build?definitionId=%d'>%s</a>`, build.ResourceContainers.Collection.BaseURL, build.Resource.Project.Name, build.Resource.Definition.ID, build.Resource.Definition.Name) + "\n"
	msg += fmt.Sprintf(`<b>Branch:</b> %s`, build.Resource.SourceBranch) + "\n"
	msg += fmt.Sprintf(`<b>Requested by:</b> %s`, build.Resource.RequestedBy.DisplayName) + "\n\n"
	msg += build.Message.HTML + "\n\n"
	msg += fmt.Sprintf(`<a href='%s%s/_build/results?buildId=%d&view=logs'>Logs</a>`, build.ResourceContainers.Collection.BaseURL, build.Resource.Project.Name, build.Resource.ID)

	tgSendMessage(msg)

	/*
		p := jsonMap.(map[string]interface{})["detailedMessage"]
		p1 := p.(map[string]interface{})["html"]
		detailedMessage := fmt.Sprintf("%v", p1)

		p = jsonMap.(map[string]interface{})["resource"]
		p1 = p.(map[string]interface{})["result"]
		buildResult := fmt.Sprintf("%v", p1)
		switch buildResult {
		case "succeeded":
			buildResult = "✅ " + buildResult
			break
		case "failed":
			buildResult = "❌ " + buildResult
			break
		case "partiallySucceeded":
			buildResult = "⚠️ " + buildResult
			break
		}

		p = jsonMap.(map[string]interface{})["resource"]
		p1 = p.(map[string]interface{})["definition"]
		p2 := p1.(map[string]interface{})["name"]
		definition := fmt.Sprintf("%v", p2)

		msg = buildResult + "\n\n" + detailedMessage + "\nDefinition: " + definition

		msg = strings.ReplaceAll(msg, "<nil>", "")
		msg = strings.ReplaceAll(msg, "</nil>", "")
		msg = strings.ReplaceAll(msg, "<ul>", "")
		msg = strings.ReplaceAll(msg, "</li>", "")
		msg = strings.ReplaceAll(msg, "</ul>", "")
		msg = strings.ReplaceAll(msg, "<li>", "⨠ ")

		log.Println(msg)

		b, err := tb.NewBot(tb.Settings{
			Token: cfg.Telegram.BotToken,
		})
		if err != nil {
			log.Println(err)
		} else {
			group := tb.ChatID(cfg.Telegram.BuildChatID)
			var opts tb.SendOptions
			opts.ParseMode = tb.ModeHTML
			e, err := b.Send(group, msg, &opts)
			log.Println(e)
			if err != nil {
				log.Println(err)
			}
		}
	*/
	ctx.JSON(http.StatusOK, gin.H{"status": "Done"})
}
