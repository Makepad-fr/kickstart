package cmd

import (
	"fmt"
	"os"

	"github.com/Makepad-fr/kickstart/internal/application"
	"github.com/Makepad-fr/kickstart/internal/chart"
	"github.com/Makepad-fr/kickstart/internal/project"
)

func Execute() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  kickstart createproject <project-name>       - Create initial project structure")
		fmt.Println("  kickstart addchart <project-name> <chart-name> - Add a new chart")
		fmt.Println("  kickstart addapp <project-name> <app-name>     - Add a new application")
		return
	}

	command := os.Args[1]

	switch command {
	case "createproject":
		if len(os.Args) != 3 {
			fmt.Println("Usage: kickstart createproject <project-name>")
			return
		}
		projectName := os.Args[2]
		if err := project.CreateProjectStructure(projectName); err != nil {
			fmt.Printf("Error creating project structure: %v\n", err)
		}

	case "addchart":
		if len(os.Args) != 4 {
			fmt.Println("Usage: kickstart addchart <project-name> <chart-name>")
			return
		}
		projectName := os.Args[2]
		chartName := os.Args[3]
		if err := chart.AddChart(projectName, chartName); err != nil {
			fmt.Printf("Error adding chart: %v\n", err)
		}

	case "addapp":
		if len(os.Args) != 4 {
			fmt.Println("Usage: kickstart addapp <project-name> <app-name>")
			return
		}
		projectName := os.Args[2]
		appName := os.Args[3]
		if err := application.AddApplication(projectName, appName); err != nil {
			fmt.Printf("Error adding application: %v\n", err)
		}

	default:
		fmt.Println("Invalid command. Use one of the following:")
		fmt.Println("  kickstart createproject <project-name>       - Create initial project structure")
		fmt.Println("  kickstart addchart <project-name> <chart-name> - Add a new chart")
		fmt.Println("  kickstart addapp <project-name> <app-name>     - Add a new application")
	}
}
