package project

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Makepad-fr/kickstart/internal/skaffold"
)

var initialProjectStructure = map[string][]string{
	".github/workflows": {},
	"charts":            {},
	"applications":      {},
	".":                 {"skaffold.yaml", ".helmignore", "README.md"},
}

func CreateProjectStructure(baseDir string) error {
	for dir, files := range initialProjectStructure {
		dirPath := filepath.Join(baseDir, dir)
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return err
		}

		for _, file := range files {
			filePath := filepath.Join(dirPath, file)
			if _, err := os.Create(filePath); err != nil {
				return err
			}
		}
	}

	// Initialize skaffold.yaml with an empty deploy section
	if err := skaffold.InitializeSkaffold(filepath.Join(baseDir, "skaffold.yaml")); err != nil {
		return err
	}

	fmt.Println("Project structure created successfully!")
	return nil
}
