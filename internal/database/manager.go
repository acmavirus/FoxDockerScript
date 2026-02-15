// Copyright by AcmaTvirus
package database

import (
	"fmt"
	"os/exec"
	"strings"
)

type DatabaseContainer struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Image   string `json:"image"`
	Status  string `json:"status"`
	Type    string `json:"type"` // mysql, postgres, redis, etc.
	Port    string `json:"port"`
}

func ListDatabases() ([]DatabaseContainer, error) {
	// Filter for containers with common database images
	cmd := exec.Command("docker", "ps", "--format", "{{.ID}}|{{.Names}}|{{.Image}}|{{.Status}}|{{.Ports}}")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	var databases []DatabaseContainer
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "|")
		if len(parts) < 5 {
			continue
		}

		image := strings.ToLower(parts[2])
		dbType := ""
		if strings.Contains(image, "mysql") {
			dbType = "MySQL"
		} else if strings.Contains(image, "postgres") {
			dbType = "PostgreSQL"
		} else if strings.Contains(image, "redis") {
			dbType = "Redis"
		} else if strings.Contains(image, "mongodb") || strings.Contains(image, "mongo") {
			dbType = "MongoDB"
		} else if strings.Contains(image, "mariadb") {
			dbType = "MariaDB"
		}

		if dbType != "" {
			databases = append(databases, DatabaseContainer{
				ID:     parts[0],
				Name:   parts[1],
				Image:  parts[2],
				Status: parts[3],
				Type:   dbType,
				Port:   parts[4],
			})
		}
	}

	return databases, nil
}

func CreateDatabase(dbType, name, password string) error {
	var image, port, env string
	switch strings.ToLower(dbType) {
	case "mysql":
		image = "mysql:8.0"
		port = "3306"
		env = "MYSQL_ROOT_PASSWORD=" + password
	case "postgres":
		image = "postgres:15"
		port = "5432"
		env = "POSTGRES_PASSWORD=" + password
	case "redis":
		image = "redis:7.0-alpine"
		port = "6379"
		// Redis usually doesn't have an ENV for password, but we'll use a simple approach here for now
	case "mongodb":
		image = "mongo:6.0"
		port = "27017"
		env = "MONGO_INITDB_ROOT_PASSWORD=" + password
	default:
		return fmt.Errorf("unsupported database type: %s", dbType)
	}

	args := []string{"run", "-d", "--name", name, "-p", port + ":" + port}
	if env != "" {
		args = append(args, "-e", env)
	}
	args = append(args, image)

	cmd := exec.Command("docker", args...)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to create db: %v, output: %s", err, string(output))
	}

	return nil
}
