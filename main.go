package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/kardianos/service"
	tb "gopkg.in/tucnak/telebot.v2"
)

const DEBUG = false

var logger service.Logger
var cfg Config

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}

func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func (p *program) run() {
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

func myMain() {
	readConfigFile(&cfg)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.HTMLRender = ginview.Default()

	router.POST("/wit/ci/create", witCiCreate)

	router.Run(":" + cfg.Web.Port)
}

func witCiCreate(ctx *gin.Context) {
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
		msg := m.(map[string]interface{})["markdown"]

		b, err := tb.NewBot(tb.Settings{
			Token: cfg.Telegram.BotToken,
		})
		if err != nil {
			log.Println(err)
		} else {
			group := tb.ChatID(cfg.Telegram.ChatID)
			b.Send(group, msg)
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
