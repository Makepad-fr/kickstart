package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Makepad-fr/kickstart/internal/application"
	"github.com/Makepad-fr/kickstart/internal/chart"
	"github.com/Makepad-fr/kickstart/internal/project"
)

func Execute() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "init-project":
		if len(os.Args) != 3 {
			fmt.Println("Usage: kickstart init-project <project-name>")
			return
		}
		projectName := os.Args[2]
		if err := project.CreateProjectStructure(projectName); err != nil {
			fmt.Printf("Error creating project structure: %v\n", err)
		}

	case "add-chart":
		var projectName, chartName string
		if len(os.Args) == 4 {
			projectName = os.Args[2]
			chartName = os.Args[3]
		} else if len(os.Args) == 3 {
			projectName = "."
			chartName = os.Args[2]
		} else {
			fmt.Println("Usage: kickstart add-chart <project-name> <chart-name> or kickstart add-chart <chart-name>")
			return
		}
		if err := chart.AddChart(filepath.Clean(projectName), chartName); err != nil {
			fmt.Printf("Error adding chart: %v\n", err)
		}

	case "add-app":
		var projectName, appName string
		if len(os.Args) == 4 {
			projectName = os.Args[2]
			appName = os.Args[3]
		} else if len(os.Args) == 3 {
			projectName = "."
			appName = os.Args[2]
		} else {
			fmt.Println("Usage: kickstart add-app <project-name> <app-name> or kickstart add-app <app-name>")
			return
		}
		if err := application.AddApplication(filepath.Clean(projectName), appName); err != nil {
			fmt.Printf("Error adding application: %v\n", err)
		}

	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  kickstart init-project <project-name>               - Create initial project structure")
	fmt.Println("  kickstart add-chart <project-name> <chart-name>     - Add a new chart (project-name optional)")
	fmt.Println("  kickstart add-chart <chart-name>                    - Add a new chart to the current directory")
	fmt.Println("  kickstart add-app <project-name> <app-name>         - Add a new application (project-name optional)")
	fmt.Println("  kickstart add-app <app-name>                        - Add a new application to the current directory")
}
