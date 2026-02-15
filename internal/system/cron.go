// Copyright by AcmaTvirus
package system

import (
	"encoding/json"
	"os"
)

type CronJob struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Schedule string `json:"schedule"` // e.g. "0 0 * * *"
	Command  string `json:"command"`
	Status   string `json:"status"`
}

const cronFile = "data/cron.json"

func GetCronJobs() ([]CronJob, error) {
	file, err := os.ReadFile(cronFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []CronJob{}, nil
		}
		return nil, err
	}
	var jobs []CronJob
	err = json.Unmarshal(file, &jobs)
	return jobs, err
}

func SaveCronJobs(jobs []CronJob) error {
	data, err := json.MarshalIndent(jobs, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(cronFile, data, 0644)
}
