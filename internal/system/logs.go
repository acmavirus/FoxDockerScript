// Copyright by AcmaTvirus
package system

import (
	"fmt"
	"os/exec"
	"strings"
)

type LogLine struct {
	Time    string `json:"time"`
	Source  string `json:"source"`
	Level   string `json:"level"`
	Message string `json:"message"`
}

func GetSystemLogs(logType string, lines int) ([]string, error) {
	var cmd *exec.Cmd

	switch logType {
	case "syslog":
		cmd = exec.Command("tail", "-n", fmt.Sprintf("%d", lines), "/var/log/syslog")
	case "auth":
		cmd = exec.Command("tail", "-n", fmt.Sprintf("%d", lines), "/var/log/auth.log")
	case "foxdocker":
		cmd = exec.Command("tail", "-n", fmt.Sprintf("%d", lines), "data/foxdocker.log")
	case "docker":
		// Get last 50 docker events
		cmd = exec.Command("docker", "events", "--since", "1h", "--until", "0s")
		// Note: docker events is blocking, tailing is better for specific containers.
		// For general events, we'll just take a snapshot
	default:
		return nil, fmt.Errorf("unknown log type: %s", logType)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return []string{fmt.Sprintf("Logs not available: %v", err)}, nil
	}

	content := string(output)
	if content == "" {
		return []string{}, nil
	}

	// Split and clean
	rawLines := strings.Split(content, "\n")
	result := []string{}
	for _, l := range rawLines {
		if l != "" {
			result = append(result, l)
		}
	}

	return result, nil
}

