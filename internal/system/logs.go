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
		filePath := "/var/log/syslog"
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			filePath = "/var/log/messages" // Fallback for CentOS/RHEL
		}
		cmd = exec.Command("tail", "-n", fmt.Sprintf("%d", lines), filePath)
	case "auth":
		filePath := "/var/log/auth.log"
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			filePath = "/var/log/secure" // Fallback for CentOS/RHEL
		}
		cmd = exec.Command("tail", "-n", fmt.Sprintf("%d", lines), filePath)
	case "foxdocker":
		cmd = exec.Command("tail", "-n", fmt.Sprintf("%d", lines), "data/foxdocker.log")
	case "docker":
		// Snapshot of recent events to avoid blocking SSE
		cmd = exec.Command("docker", "events", "--since", "5m", "--until", "0s")
	default:
		return nil, fmt.Errorf("unknown log type: %s", logType)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return []string{fmt.Sprintf("Logs not accessible: %v. Please ensure /var/log is mounted and the file exists.", err)}, nil
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

