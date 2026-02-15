// Copyright by AcmaTvirus
package system

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// App represents an application template
type App struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Image string   `json:"image"`
	Ports string   `json:"ports"`
	Env   []string `json:"env"`
}

// InstallApp handles the installation of an app using Docker Compose
func InstallApp(app App, envVars map[string]string) error {
	appDir := filepath.Join("/opt/foxdocker/apps", app.ID)
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return fmt.Errorf("failed to create app directory: %v", err)
	}

	composeContent := generateComposeFile(app, envVars)
	composePath := filepath.Join(appDir, "docker-compose.yml")
	if err := os.WriteFile(composePath, []byte(composeContent), 0644); err != nil {
		return fmt.Errorf("failed to write docker-compose.yml: %v", err)
	}

	// Execute docker compose up -d
	cmd := exec.Command("docker", "compose", "-f", composePath, "up", "-d")
	cmd.Dir = appDir
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("docker compose failed: %v, output: %s", err, string(output))
	}

	return nil
}

func generateComposeFile(app App, envVars map[string]string) string {
	var sb strings.Builder
	sb.WriteString("version: '3.8'\n")
	sb.WriteString("services:\n")
	sb.WriteString(fmt.Sprintf("  %s:\n", app.ID))
	sb.WriteString(fmt.Sprintf("    image: %s\n", app.Image))
	
	if app.Ports != "" {
		sb.WriteString("    ports:\n")
		// Simplified port mapping: mapping single port to itself
		sb.WriteString(fmt.Sprintf("      - \"%s:%s\"\n", app.Ports, app.Ports))
	}

	if len(app.Env) > 0 {
		sb.WriteString("    environment:\n")
		for _, key := range app.Env {
			val := envVars[key]
			sb.WriteString(fmt.Sprintf("      - %s=%s\n", key, val))
		}
	}

	sb.WriteString("    restart: always\n")
	sb.WriteString("networks:\n")
	sb.WriteString("  default:\n")
	sb.WriteString("    name: foxdocker-network\n")
	sb.WriteString("    external: true\n")

	return sb.String()
}
