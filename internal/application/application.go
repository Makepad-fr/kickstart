package application

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Makepad-fr/kickstart/internal/skaffold"
)

func AddApplication(baseDir, appName string) error {
	appDir := filepath.Join(baseDir, "applications", appName)

	if err := os.MkdirAll(appDir, os.ModePerm); err != nil {
		return err
	}

	// Create Go module files
	if err := createGoModFiles(appDir, appName); err != nil {
		return err
	}

	// Create the main.go file with a basic REST API server
	if err := createMainGoFile(filepath.Join(appDir, "main.go")); err != nil {
		return err
	}

	// Create a minimal Dockerfile
	if err := createDockerfile(filepath.Join(appDir, "Dockerfile")); err != nil {
		return err
	}

	// Update skaffold.yaml
	if err := skaffold.UpdateSkaffoldForApp(filepath.Join(baseDir, "skaffold.yaml"), appName); err != nil {
		return err
	}

	fmt.Printf("Application '%s' added successfully!\n", appName)
	return nil
}

func createGoModFiles(appDir, appName string) error {
	goModPath := filepath.Join(appDir, "go.mod")
	goModContent := fmt.Sprintf(`module %s

go 1.20
`, appName)

	if err := os.WriteFile(goModPath, []byte(goModContent), 0644); err != nil {
		return err
	}

	// You may skip go.sum creation as it will be automatically generated when you run `go mod tidy` or build the application.
	return nil
}

func createMainGoFile(mainGoPath string) error {
	const mainGoTemplate = `package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})

	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
`

	// Write the main.go file
	return os.WriteFile(mainGoPath, []byte(mainGoTemplate), 0644)
}

func createDockerfile(dockerfilePath string) error {
	const dockerfileTemplate = `# Start with the official Golang image
FROM golang:1.20-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker's layer caching
COPY go.mod ./
# Run 'go mod download' to install dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application binary
RUN go build -o /go-rest-api .

# Start a new stage with a minimal Alpine Linux image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the previous stage
COPY --from=builder /go-rest-api .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./go-rest-api"]`

	// Write the Dockerfile
	return os.WriteFile(dockerfilePath, []byte(dockerfileTemplate), 0644)
}
