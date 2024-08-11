package chart

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/Makepad-fr/kickstart/internal/skaffold"
)

// ChartMetadata represents the basic structure of a Helm Chart.yaml file
type ChartMetadata struct {
	APIVersion  string
	Name        string
	Version     string
	Description string
}

func AddChart(baseDir, chartName string) error {
	chartDir := filepath.Join(baseDir, "charts", chartName)
	subDirs := []string{"charts", "templates"}
	files := []string{"values.yaml", "README.md"}

	// Create directories
	for _, subDir := range subDirs {
		if err := os.MkdirAll(filepath.Join(chartDir, subDir), os.ModePerm); err != nil {
			return err
		}
	}

	// Create files
	for _, file := range files {
		filePath := filepath.Join(chartDir, file)
		if _, err := os.Create(filePath); err != nil {
			return err
		}
	}

	// Create and populate the Chart.yaml file
	chartMetadata := ChartMetadata{
		APIVersion:  "v2", // Default API version for Helm 3
		Name:        chartName,
		Version:     "0.1.0", // Default initial version
		Description: fmt.Sprintf("A Helm chart for %s", chartName),
	}

	chartYamlPath := filepath.Join(chartDir, "Chart.yaml")
	if err := createChartYaml(chartYamlPath, chartMetadata); err != nil {
		return err
	}

	// Update skaffold.yaml
	if err := skaffold.UpdateSkaffoldForChart(filepath.Join(baseDir, "skaffold.yaml"), chartName); err != nil {
		return err
	}

	fmt.Printf("Chart '%s' added successfully!\n", chartName)
	return nil
}

func createChartYaml(chartYamlPath string, metadata ChartMetadata) error {
	// Define the template for the Chart.yaml content
	const chartYamlTemplate = `apiVersion: {{.APIVersion}}
name: {{.Name}}
version: {{.Version}}
description: {{.Description}}
`

	// Parse the template
	tmpl, err := template.New("chartYaml").Parse(chartYamlTemplate)
	if err != nil {
		return err
	}

	// Create or open the Chart.yaml file
	file, err := os.Create(chartYamlPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Execute the template with the metadata and write to the file
	if err := tmpl.Execute(file, metadata); err != nil {
		return err
	}

	return nil
}
