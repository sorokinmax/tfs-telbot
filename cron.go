package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func CheckOpenTasks() {
	log.Println("CheckOpenTasks job stared")
	client := http.Client{}
	url := fmt.Sprintf("%s/_apis/wit/wiql/%s?api-version=6.0", normalizeUrl(cfg.Tfs.ProjectUrl), cfg.Tfs.AssessmentTasksQueue)
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth("", cfg.Tfs.Pat)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	query := Query{}
	jsonErr := json.Unmarshal(body, &query)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	if len(query.WorkItemRelations) > 0 {
		msg := fmt.Sprintf("Есть <b><a href='%s/_queries/query/%s/'>незакрытые задачи</a></b> на предварительную оценку.\nПросьба лидов команд проверить свои задачи ASAP.\n\n", normalizeUrl(cfg.Tfs.ProjectUrl), cfg.Tfs.AssessmentTasksQueue)
		tgSendMessage(msg, cfg.Telegram.AssessmentTasksChatID)
	}
}

func normalizeUrl(url string) string {
	return strings.ToLower(strings.TrimRight(url, "/"))
}
