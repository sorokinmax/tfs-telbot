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

const DEBUG = false

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
	router.POST("/build/report", tfsBuild)

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
	if DEBUG {
		myMain()
	} else {
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

		m := jsonMap.(map[string]interface{})["detailedMessage"]
		msg := m.(map[string]interface{})["html"]

		b, err := tb.NewBot(tb.Settings{
			Token: cfg.Telegram.BotToken,
		})
		if err != nil {
			log.Println(err)
		} else {
			group := tb.ChatID(cfg.Telegram.WitCiCreateChatID)
			_, err = b.Send(group, msg)
			if err != nil {
				log.Println(err)
			}
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Done"})
}

func tfsBuild(ctx *gin.Context) {
	var jsonMap interface{}

	data, _ := ctx.GetRawData()
	err := json.Unmarshal(data, &jsonMap)
	if err != nil {
		logger.Error(err)
	}

	//logger.Info(jsonMap)

	p := jsonMap.(map[string]interface{})["detailedMessage"]
	p1 := p.(map[string]interface{})["html"]
	msg := fmt.Sprintf("%v", p1)
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
		buildResult += " ✅"
		break
	case "failed":
		buildResult += " ❌"
		break
	case "partiallySucceeded":
		buildResult += " ⚠️"
		break
	}

	p = jsonMap.(map[string]interface{})["resource"]
	p1 = p.(map[string]interface{})["definition"]
	p2 := p1.(map[string]interface{})["name"]
	definition := fmt.Sprintf("%v", p2)

	msg += "\nDefinition: " + definition + "\nBuild result: " + buildResult

	msg = strings.ReplaceAll(msg, "<ul>", "")
	msg = strings.ReplaceAll(msg, "</li>", "")
	msg = strings.ReplaceAll(msg, "</ul>", "")
	msg = strings.ReplaceAll(msg, "<li>", "⨠ ")

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

	ctx.JSON(http.StatusOK, gin.H{"status": "Done"})
}

func currentDir() string {
	if !DEBUG {
		fullPath, err := os.Executable()
		if err != nil {
			panic(err)
		}
		path := filepath.Dir(fullPath)
		return path
	}
	return "."
}
