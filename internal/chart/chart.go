package chart

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Makepad-fr/kickstart/internal/skaffold"
)

func AddChart(baseDir, chartName string) error {
	chartDir := filepath.Join(baseDir, "charts", chartName)
	subDirs := []string{"charts", "templates"}
	files := []string{"Chart.yaml", "values.yaml", "README.md"}

	for _, subDir := range subDirs {
		if err := os.MkdirAll(filepath.Join(chartDir, subDir), os.ModePerm); err != nil {
			return err
		}
	}

	for _, file := range files {
		filePath := filepath.Join(chartDir, file)
		if _, err := os.Create(filePath); err != nil {
			return err
		}
	}

	// Update skaffold.yaml
	if err := skaffold.UpdateSkaffoldForChart(filepath.Join(baseDir, "skaffold.yaml"), chartName); err != nil {
		return err
	}

	fmt.Printf("Chart '%s' added successfully!\n", chartName)
	return nil
}
