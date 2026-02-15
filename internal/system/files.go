// Copyright by AcmaTvirus
package system

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileItem struct {
	Name  string `json:"name"`
	Path  string `json:"path"`
	IsDir bool   `json:"is_dir"`
	Size  int64  `json:"size"`
	ModTime string `json:"mod_time"`
}

const ProjectsRoot = "/opt/foxdocker/apps"
const BackupRoot = "/opt/foxdocker/backups"

func ListFiles(path string) ([]FileItem, error) {
	fullPath := filepath.Join(ProjectsRoot, path)
	// Security check: ensure path is within ProjectsRoot
	if !filepath.HasPrefix(fullPath, ProjectsRoot) {
		return nil, fmt.Errorf("access denied")
	}

	entries, err := os.ReadDir(fullPath)
	if err != nil {
		return nil, err
	}

	var items []FileItem
	for _, entry := range entries {
		info, _ := entry.Info()
		items = append(items, FileItem{
			Name:    entry.Name(),
			Path:    filepath.Join(path, entry.Name()),
			IsDir:   entry.IsDir(),
			Size:    info.Size(),
			ModTime: info.ModTime().Format("2006-01-02 15:04:05"),
		})
	}

	return items, nil
}

func ReadFileContent(path string) (string, error) {
	fullPath := filepath.Join(ProjectsRoot, path)
	if !filepath.HasPrefix(fullPath, ProjectsRoot) {
		return "", fmt.Errorf("access denied")
	}

	content, err := os.ReadFile(fullPath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func SaveFileContent(path, content string) error {
	fullPath := filepath.Join(ProjectsRoot, path)
	if !filepath.HasPrefix(fullPath, ProjectsRoot) {
		return fmt.Errorf("access denied")
	}

	return os.WriteFile(fullPath, []byte(content), 0644)
}
