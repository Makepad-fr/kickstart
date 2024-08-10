package skaffold

import (
	"os"
	"path/filepath"

	"sigs.k8s.io/yaml"
)

type SkaffoldConfig struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Deploy     Deploy `yaml:"deploy"`
}

type Deploy struct {
	Helm Helm `yaml:"helm"`
}

type Helm struct {
	Releases []Release `yaml:"releases"`
}

type Release struct {
	Name      string `yaml:"name"`
	ChartPath string `yaml:"chartPath"`
}

func InitializeSkaffold(skaffoldPath string) error {
	initialContent := SkaffoldConfig{
		APIVersion: "skaffold/v2beta26",
		Kind:       "Config",
		Deploy: Deploy{
			Helm: Helm{
				Releases: []Release{},
			},
		},
	}

	data, err := yaml.Marshal(&initialContent)
	if err != nil {
		return err
	}

	return os.WriteFile(skaffoldPath, data, 0644)
}

func UpdateSkaffoldForChart(skaffoldPath, chartName string) error {
	data, err := os.ReadFile(skaffoldPath)
	if err != nil {
		return err
	}

	var config SkaffoldConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return err
	}

	// Add the new chart to the releases
	newRelease := Release{
		Name:      chartName,
		ChartPath: filepath.Join("charts", chartName),
	}
	config.Deploy.Helm.Releases = append(config.Deploy.Helm.Releases, newRelease)

	// Marshal back to YAML
	updatedData, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}

	return os.WriteFile(skaffoldPath, updatedData, 0644)
}

func UpdateSkaffoldForApp(skaffoldPath, appName string) error {
	data, err := os.ReadFile(skaffoldPath)
	if err != nil {
		return err
	}

	var config SkaffoldConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return err
	}

	// Assuming we want to add something specific for the application to skaffold.yaml
	// Here we add an application like a Helm release
	newRelease := Release{
		Name:      appName,
		ChartPath: filepath.Join("applications", appName),
	}
	config.Deploy.Helm.Releases = append(config.Deploy.Helm.Releases, newRelease)

	// Marshal back to YAML
	updatedData, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}

	return os.WriteFile(skaffoldPath, updatedData, 0644)
}
