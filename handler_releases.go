package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

func tfsReleaseComplete(ctx *gin.Context) {
	var (
		release Release
		msg     string
	)

	err := ctx.BindJSON(&release)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "json binding error"})
		return
	}

	// 200 because TFS fails the webhook at 4xx codes and disables it
	if slices.Contains(cfg.ExcludedPipelines.Release, release.Resource.Environment.ReleaseDefinition.Name) {
		ctx.JSON(http.StatusOK, gin.H{"Response": fmt.Sprintf(`Pipeline "%s" is exluded`, release.Resource.Environment.ReleaseDefinition.Name)})
		return
	}

	switch release.Resource.Deployment.DeploymentStatus {
	case "succeeded":
		msg += "✅✅✅\n♿️♿️♿️\n\n"
	case "failed":
		msg += "❌❌❌\n♿️♿️♿️\n\n"
	case "partiallySucceeded":
		msg += "⚠️⚠️⚠️\n♿️♿️♿️\n\n"
	case "canceled":
		msg += "✖️✖️✖️\n♿️♿️♿️\n\n"
	}

	msg += fmt.Sprintf(`<b>Definition:</b> <a href='%s%s/_release?definitionId=%d'>%s</a>`, release.ResourceContainers.Collection.BaseURL, release.Resource.Project.Name, release.Resource.Environment.ReleaseDefinition.ID, release.Resource.Environment.ReleaseDefinition.Name) + "\n"
	msg += fmt.Sprintf(`<b>Created by:</b> %s`, release.Resource.Deployment.RequestedBy.DisplayName) + "\n\n"

	msg += "<b>Last commtis:</b>\n"
	for i, commit := range release.Resource.Data.Commits {
		if i == 5 {
			break
		}
		msg += fmt.Sprintf(`- <a href='%s'>%s</a>`, commit.DisplayURI, commit.Message) + "\n"
	}
	msg += "\n"

	msg += release.Message.HTML + "\n\n"

	for _, deployStep := range release.Resource.Environment.DeploySteps {
		for _, deployPhase := range deployStep.ReleaseDeployPhases {
			for _, deploymentJob := range deployPhase.DeploymentJobs {
				for _, task := range deploymentJob.Tasks {
					if task.Status != "succeeded" {
						msg += fmt.Sprintf(`<b>Failed step:</b> %s`, task.Name) + "\n"
					}
				}
			}
		}
	}

	msg += fmt.Sprintf(`<a href='%s%s/_releaseProgress?_a=release-environment-logs&releaseId=%d'>Logs</a>`, release.ResourceContainers.Collection.BaseURL, release.Resource.Project.Name, release.Resource.Environment.ReleaseID)

	tgSendMessage(msg, cfg.Telegram.DeployChatID)

	ctx.JSON(http.StatusOK, gin.H{"status": "Done"})
}
