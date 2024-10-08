package skaffold

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"gopkg.in/yaml.v3"
)

type SkaffoldConfig struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Build      Build  `yaml:"build"`
	Deploy     Deploy `yaml:"deploy"`
}

type Build struct {
	Local     Local      `yaml:"local"`
	Artifacts []Artifact `yaml:"artifacts"`
}

type Local struct {
	Push        bool `yaml:"push"`
	Concurrency int  `yaml:"concurrency"`
}

type Artifact struct {
	ImageName string `yaml:"image"`
	Context   string `yaml:"context"`
	Docker    Docker `yaml:"docker"`
}

type Docker struct {
	Dockerfile string `yaml:"dockerfile"`
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

// GitHubFileContent represents the content of a file from the GitHub API response
type GitHubFileContent struct {
	Content  string `json:"content"`
	Encoding string `json:"encoding"`
}

var latestSkaffoldVersion string

func getLatestSkaffoldYamlVersion() (string, error) {
	// GitHub API URL for the specific file
	url := "https://api.github.com/repos/GoogleContainerTools/skaffold/contents/docs-v2/content/en/docs/references/cli/_index.md"

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error making the HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: unable to fetch the file. Status code: %d", resp.StatusCode)
	}

	// Decode the JSON response
	var fileContent GitHubFileContent
	if err := json.NewDecoder(resp.Body).Decode(&fileContent); err != nil {
		return "", fmt.Errorf("error decoding the JSON response: %w", err)
	}

	// Decode the base64-encoded content
	contentBytes, err := base64.StdEncoding.DecodeString(fileContent.Content)
	if err != nil {
		return "", fmt.Errorf("error decoding base64 content: %w", err)
	}
	content := string(contentBytes)

	// Define the regular expression to extract the version
	re := regexp.MustCompile(`--version='skaffold/(v[0-9a-zA-Z]+)'`)

	// Find the version in the file content
	matches := re.FindStringSubmatch(content)
	if len(matches) > 1 {
		return matches[1], nil
	} else {
		return "", fmt.Errorf("version not found in the file")
	}
}

func InitializeSkaffold(skaffoldPath string) error {
	skaffoldYamlVersion, err := getLatestSkaffoldYamlVersion()
	if err != nil {
		log.Printf("Error while getting latest version from GitHub: %v", err)
		skaffoldYamlVersion = latestSkaffoldVersion
	}
	initialContent := SkaffoldConfig{
		APIVersion: fmt.Sprintf("skaffold/%s", skaffoldYamlVersion),
		Kind:       "Config",
		Build: Build{
			Local: Local{
				Push:        false,
				Concurrency: 0,
			},
			Artifacts: []Artifact{},
		},
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

	// Add the new Docker artifact without specifying a tag
	newArtifact := Artifact{
		ImageName: appName, // Do not include a tag like ":latest"
		Context:   filepath.Join("applications", appName),
		Docker: Docker{
			Dockerfile: "Dockerfile",
		},
	}
	config.Build.Artifacts = append(config.Build.Artifacts, newArtifact)

	// Marshal back to YAML
	updatedData, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}

	return os.WriteFile(skaffoldPath, updatedData, 0644)
}
