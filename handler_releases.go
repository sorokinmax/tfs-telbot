package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

func tfsReleaseBegin(ctx *gin.Context) {
	var (
		release Release
		msg     string
	)

	err := ctx.BindJSON(&release)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "json binding error"})
		return
	}

	if slices.Contains(cfg.ExcludedPipelines.Release, release.Resource.Environment.ReleaseDefinition.Name) {
		ctx.AbortWithStatusJSON(
			http.StatusPreconditionFailed,
			gin.H{"Response": fmt.Sprintf(`Pipeline "%s" is exluded`, release.Resource.Environment.ReleaseDefinition.Name)},
		)
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
	//msg += fmt.Sprintf(`<b>Target server:</b> %s`, release.Resource.Environment.Name) + "\n"
	//msg += fmt.Sprintf(`<b>Release name:</b> %s`, release.Resource.SourceBranch) + "\n"
	//msg += fmt.Sprintf(`<b>Build name:</b> %s`, release.Resource.SourceBranch) + "\n"
	msg += fmt.Sprintf(`<b>Created by:</b> %s`, release.Resource.Deployment.RequestedBy.DisplayName) + "\n\n"
	msg += release.Message.HTML + "\n\n"
	msg += fmt.Sprintf(`<a href='%s%s/_releaseProgress?_a=release-environment-logs&releaseId=%d'>Logs</a>`, release.ResourceContainers.Collection.BaseURL, release.Resource.Project.Name, release.Resource.Environment.ReleaseID)

	tgSendMessage(msg)

	ctx.JSON(http.StatusOK, gin.H{"status": "Done"})
}
