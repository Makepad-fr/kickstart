package application

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Makepad-fr/kickstart/internal/skaffold"
)

func AddApplication(baseDir, appName string) error {
	appDir := filepath.Join(baseDir, "applications", appName)
	files := []string{"Dockerfile", "main.go", "go.mod", "go.sum"}

	if err := os.MkdirAll(appDir, os.ModePerm); err != nil {
		return err
	}

	for _, file := range files {
		filePath := filepath.Join(appDir, file)
		if _, err := os.Create(filePath); err != nil {
			return err
		}
	}

	// Update skaffold.yaml
	if err := skaffold.UpdateSkaffoldForApp(filepath.Join(baseDir, "skaffold.yaml"), appName); err != nil {
		return err
	}

	fmt.Printf("Application '%s' added successfully!\n", appName)
	return nil
}
