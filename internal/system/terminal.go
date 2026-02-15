// Copyright by AcmaTvirus
package system

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func ExecuteContainerCommand(containerID, command string) (string, error) {
	// Execute command inside docker container
	cmd := exec.Command("docker", "exec", containerID, "sh", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), fmt.Errorf("command failed: %v", err)
	}
	return string(output), nil
}

func CreateBackup(projectID string) (string, error) {
	os.MkdirAll(BackupRoot, 0755)
	
	fileName := fmt.Sprintf("%s_%d.tar.gz", projectID, time.Now().Unix())
	filePath := filepath.Join(BackupRoot, fileName)
	projectPath := filepath.Join(ProjectsRoot, projectID)
	
	cmd := exec.Command("tar", "-czf", filePath, "-C", projectPath, ".")
	if output, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("backup failed: %v, output: %s", err, string(output))
	}
	
	return filePath, nil
}
