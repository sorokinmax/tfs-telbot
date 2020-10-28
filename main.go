package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/kardianos/service"
	tb "gopkg.in/tucnak/telebot.v2"
)

var logger service.Logger
var cfg Config

type program struct{}

func myMain() {
	gin.SetMode(gin.ReleaseMode)
	f, _ := os.Create(currentDir() + "/web.log")
	gin.DefaultWriter = io.MultiWriter(os.Stdout, f)
	defer f.Close()

	readConfigFile(&cfg)

	router := gin.Default()
	router.HTMLRender = ginview.Default()

	router.POST("/wit/ci/create", tfsWitCiCreate)
	router.POST("/build/report", tfsBuildReport)
	router.POST("/release/begin", tfsReleaseBegin)

	router.Run(":" + cfg.Web.Port)
}

func (p *program) Start(s service.Service) error {
	if service.Interactive() {
		logger.Info("Running in terminal.")
		myMain()
	} else {
		logger.Info("Running under service manager.")
	}

	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}

func (p *program) Stop(s service.Service) error {
	logger.Info("I'm Stopping")
	// Stop should not block. Return with a few seconds.
	return nil
}

func (p *program) run() {
	logger.Infof("I'm running %v.", service.Platform())
	myMain()
}

func main() {

	svcConfig := &service.Config{
		Name:        "TFS-GoBot",
		DisplayName: "TFS GoBot",
		Description: "TFS notifier telegram bot",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) > 1 {
		err = service.Control(s, os.Args[1])
		if err != nil {
			log.Printf("Valid actions: %q\n", service.ControlAction)
			log.Fatal(err)
		}
		return
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}

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
				logger.Error(err)
			}
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Done"})
}

func tfsBuildReport(ctx *gin.Context) {
	var jsonMap interface{}
	var msg = ""

	data, _ := ctx.GetRawData()
	err := json.Unmarshal(data, &jsonMap)
	if err != nil {
		logger.Error(err)
	}

	//logger.Info(jsonMap)

	p := jsonMap.(map[string]interface{})["detailedMessage"]
	p1 := p.(map[string]interface{})["html"]
	detailedMessage := fmt.Sprintf("%v", p1)
	/*
		p = jsonMap.(map[string]interface{})["resource"]
		p1 = p.(map[string]interface{})["buildNumber"]
		buildNumber := fmt.Sprintf("%v", p1)
	*/
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

	logger.Info(msg)

	b, err := tb.NewBot(tb.Settings{
		Token: cfg.Telegram.BotToken,
	})
	if err != nil {
		logger.Error(err)
	} else {
		group := tb.ChatID(cfg.Telegram.BuildChatID)
		var opts tb.SendOptions
		opts.ParseMode = tb.ModeHTML
		e, err := b.Send(group, msg, &opts)
		log.Println(e)
		if err != nil {
			logger.Error(err)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Done"})
}

func tfsReleaseBegin(ctx *gin.Context) {
	var jsonMap interface{}
	var msg = ""

	data, _ := ctx.GetRawData()
	err := json.Unmarshal(data, &jsonMap)
	if err != nil {
		logger.Error(err)
	}

	//logger.Info(jsonMap)
	logger.Info(string(data))

	p := jsonMap.(map[string]interface{})["detailedMessage"]
	p1 := p.(map[string]interface{})["html"]
	detailedMessage := fmt.Sprintf("%v", p1)

	if strings.Contains(strings.ToLower(detailedMessage), strings.ToLower(">DEV<")) {

		/*
			p = jsonMap.(map[string]interface{})["resource"]
			p1 = p.(map[string]interface{})["environment"]
			p2 := p1.(map[string]interface{})["name"]
			envName := fmt.Sprintf("%v", p2)
		*/
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

		/*
			p = jsonMap.(map[string]interface{})["resource"]
			p1 = p.(map[string]interface{})["environment"]
			p2 = p1.(map[string]interface{})["scheduledDeploymentTime"]
			scheduledDeploymentTime := fmt.Sprintf("%v", p2)
		*/

		msg = "♿️ deploying\n" + detailedMessage + "\n" + "\nTarget server: " + targetServer + "\nRelease name: " + releaseName + "\nBuild name: " + buildName + "\nCreated by: " + createdBy + "\n\nTime to deploy: " + timeToDeploy

		msg = strings.ReplaceAll(msg, "<nil>", "nil")
		msg = strings.ReplaceAll(msg, "</nil>", "")
		msg = strings.ReplaceAll(msg, "<br>", "\n")
		msg = strings.ReplaceAll(msg, "<ul>", "")
		msg = strings.ReplaceAll(msg, "</li>", "")
		msg = strings.ReplaceAll(msg, "</ul>", "")
		msg = strings.ReplaceAll(msg, "<li>", "⨠ ")

		logger.Info(msg)
		//log.Println(msg)

		b, err := tb.NewBot(tb.Settings{
			Token: cfg.Telegram.BotToken,
		})
		if err != nil {
			logger.Error(err)
		} else {
			group := tb.ChatID(cfg.Telegram.BuildChatID)
			var opts tb.SendOptions
			opts.ParseMode = tb.ModeHTML
			e, err := b.Send(group, msg, &opts)
			log.Println(e)
			if err != nil {
				logger.Error(err)
			}
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "Done"})
}

func currentDir() string {
	fullPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	path := filepath.Dir(fullPath)
	return path
}
