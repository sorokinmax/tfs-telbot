package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	//"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	tb "gopkg.in/telebot.v3"
)

var cfg Config

func main() {

	log.SetFlags(log.LstdFlags)
	lf, err := os.OpenFile("telbot.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	multi := io.MultiWriter(os.Stdout, lf)
	log.SetOutput(multi)
	defer lf.Close()

	gin.SetMode(gin.ReleaseMode)

	alf, err := os.OpenFile("access.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	gin.DefaultWriter = io.MultiWriter(os.Stdout, alf)
	defer alf.Close()

	cron := gocron.NewScheduler(time.Local)
	cron.Every(1).Minute().Do(readConfigFile, &cfg)
	cron.StartAsync()
	readConfigFile(&cfg)

	router := gin.Default()
	router.Use(ginBodyLogMiddleware)

	router.POST("/wit/ci/create", tfsWitCiCreate)
	router.POST("/build/report", tfsBuildReport)
	router.POST("/release/begin", tfsReleaseBegin)

	router.Run(":" + cfg.Web.Port)
}

/*
	func tgPrepareMD2String(msg string) string {
		r := strings.NewReplacer(
			"_", "\\_",
			"*", "\\*",
			"[", "\\[",
			"]", "\\]",
			"(", "\\(",
			")", "\\)",
			"~", "\\~",
			"`", "\\`",
			">", "\\>",
			"#", "\\#",
			"+", "\\+",
			"-", "\\-",
			"=", "\\=",
			"|", "\\|",
			"{", "\\{",
			"}", "\\}",
			".", "\\.",
			"!", "\\!",
		)
		return r.Replace(msg)
	}
*/
func tfsWitCiCreate(ctx *gin.Context) {
	var jsonMap interface{}

	data, _ := ctx.GetRawData()
	err := json.Unmarshal(data, &jsonMap)
	if err != nil {
		log.Println(err)
	}

	p := jsonMap.(map[string]interface{})["resource"]
	p1 := p.(map[string]interface{})["fields"]
	p2 := p1.(map[string]interface{})["Microsoft.VSTS.Common.Priority"]
	priority := fmt.Sprintf("%v", p2)

	if priority == "1" {

		p := jsonMap.(map[string]interface{})["message"]
		p1 := p.(map[string]interface{})["html"]
		msg := fmt.Sprintf("%v", p1)

		p = jsonMap.(map[string]interface{})["resource"]
		p1 = p.(map[string]interface{})["fields"]
		p2 := p1.(map[string]interface{})["bm.Client"]
		client := fmt.Sprintf("%v", p2)

		p = jsonMap.(map[string]interface{})["resource"]
		p1 = p.(map[string]interface{})["fields"]
		p2 = p1.(map[string]interface{})["bm.ServerVersion"]
		serverVersion := fmt.Sprintf("%v", p2)

		p = jsonMap.(map[string]interface{})["resource"]
		p1 = p.(map[string]interface{})["fields"]
		p2 = p1.(map[string]interface{})["bm.IpadVersion"]
		ipadVersion := fmt.Sprintf("%v", p2)

		p = jsonMap.(map[string]interface{})["resource"]
		p1 = p.(map[string]interface{})["fields"]
		p2 = p1.(map[string]interface{})["bm.Zendesk"]
		zendesk := fmt.Sprintf("%v", p2)

		msg += "\n\nClient: " + client + "\nServer version: " + serverVersion + "\niPad version: " + ipadVersion + "\nZendesk: " + zendesk

		msg = strings.ReplaceAll(msg, "<ul>", "")
		msg = strings.ReplaceAll(msg, "</li>", "")
		msg = strings.ReplaceAll(msg, "</ul>", "")
		msg = strings.ReplaceAll(msg, "<li>", "⨠ ")

		b, err := tb.NewBot(tb.Settings{
			Token: cfg.Telegram.BotToken,
		})
		if err != nil {
			log.Println(err)
		} else {
			group := tb.ChatID(cfg.Telegram.WitCiCreateChatID)
			var opts tb.SendOptions
			opts.ParseMode = tb.ModeHTML
			e, err := b.Send(group, msg, &opts)
			log.Println(e)
			if err != nil {
				log.Println(err)
			}
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Done"})
}

/*
func tfsReleaseBegin(ctx *gin.Context) {
	var jsonMap interface{}
	var msg = ""

	data, _ := ctx.GetRawData()
	err := json.Unmarshal(data, &jsonMap)
	if err != nil {
		log.Println(err)
	}

	//logger.Info(jsonMap)
	log.Println(string(data))

	p := jsonMap.(map[string]interface{})["detailedMessage"]
	p1 := p.(map[string]interface{})["html"]
	detailedMessage := fmt.Sprintf("%v", p1)

	if strings.Contains(strings.ToLower(detailedMessage), strings.ToLower(">DEV<")) {


		//	p = jsonMap.(map[string]interface{})["resource"]
		//	p1 = p.(map[string]interface{})["environment"]
		//	p2 := p1.(map[string]interface{})["name"]
		//	envName := fmt.Sprintf("%v", p2)

		p = jsonMap.(map[string]interface{})["resource"]
		p1 = p.(map[string]interface{})["environment"]
		p2 := p1.(map[string]interface{})["variables"]
		p3 := p2.(map[string]interface{})["TargetServer"]
		p4 := p3.(map[string]interface{})["value"]
		targetServer := fmt.Sprintf("%v", p4)

		p = jsonMap.(map[string]interface{})["resource"]
		p1 = p.(map[string]interface{})["environment"]
		p2 = p1.(map[string]interface{})["timeToDeploy"]
		timeToDeploy := fmt.Sprintf("%v", p2)

		p = jsonMap.(map[string]interface{})["resource"]
		p1 = p.(map[string]interface{})["release"]
		p2 = p1.(map[string]interface{})["name"]
		releaseName := fmt.Sprintf("%v", p2)

		p = jsonMap.(map[string]interface{})["resource"]
		p1 = p.(map[string]interface{})["release"]
		p2 = p1.(map[string]interface{})["createdBy"]
		p3 = p2.(map[string]interface{})["displayName"]
		createdBy := fmt.Sprintf("%v", p3)

		p = jsonMap.(map[string]interface{})["resource"]
		p1 = p.(map[string]interface{})["release"]
		p2 = p1.(map[string]interface{})["artifacts"].([]interface{})[0]
		p3 = p2.(map[string]interface{})["definitionReference"]
		p4 = p3.(map[string]interface{})["version"]
		p5 := p4.(map[string]interface{})["name"]
		buildName := fmt.Sprintf("%v", p5)

		//	p = jsonMap.(map[string]interface{})["resource"]
		//	p1 = p.(map[string]interface{})["environment"]
		//	p2 = p1.(map[string]interface{})["scheduledDeploymentTime"]
		//	scheduledDeploymentTime := fmt.Sprintf("%v", p2)

		msg = "♿️ deploying\n" + detailedMessage + "\n" + "\nTarget server: " + targetServer + "\nRelease name: " + releaseName + "\nBuild name: " + buildName + "\nCreated by: " + createdBy + "\n\nTime to deploy: " + timeToDeploy

		msg = strings.ReplaceAll(msg, "<nil>", "nil")
		msg = strings.ReplaceAll(msg, "</nil>", "")
		msg = strings.ReplaceAll(msg, "<br>", "\n")
		msg = strings.ReplaceAll(msg, "<ul>", "")
		msg = strings.ReplaceAll(msg, "</li>", "")
		msg = strings.ReplaceAll(msg, "</ul>", "")
		msg = strings.ReplaceAll(msg, "<li>", "⨠ ")

		log.Println(msg)
		//log.Println(msg)

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
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "Done"})
}
*/
