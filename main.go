package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
)

var cfg Config

func main() {

	// common log file
	log.SetFlags(log.LstdFlags)
	lf, err := os.OpenFile("./telbot.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	multi := io.MultiWriter(os.Stdout, lf)
	log.SetOutput(multi)
	defer lf.Close()

	gin.SetMode(gin.ReleaseMode)

	// access log file
	alf, err := os.OpenFile("./access.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	gin.DefaultWriter = io.MultiWriter(os.Stdout, alf)
	defer alf.Close()

	// regular config reading
	cron := gocron.NewScheduler(time.Local)
	cron.Every(1).Minute().Do(readConfigFile, &cfg)
	cron.Every(1).Day().At(cfg.Tfs.AssessmentTasksDailyNotificationTime).Do(CheckOpenTasks)
	cron.StartAsync()

	readConfigFile(&cfg)
	//CheckOpenTasks()

	router := gin.Default()

	// request dumper middleware for debug mode
	router.Use(ginBodyLogMiddleware)

	// routes
	router.POST("/pr/create", tfsPRCreated)
	router.POST("/cr/create", tfsCRCreated)
	router.POST("/ci/create", tfsCICreated)
	router.POST("/ci/update", tfsCIUpdated)
	router.POST("/build/report", tfsBuildReport)
	router.POST("/release/begin", tfsReleaseBegin)

	router.Run(":" + cfg.Web.Port)
}
