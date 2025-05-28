package createService

import (
	"fmt"
	"os"
	"path/filepath"
)

func createNodeService(serviceName string) {
	serviceDir := filepath.Join("../services/", fmt.Sprintf("%s-service", serviceName))
	folders := []string{
		"src/controllers",
		"src/services",
		"src/repositories",
		"src/models",
		"src/types",
		"src/routes",
		"logger",
		"monitor",
		"tests",
	}

	fmt.Println("📦 Creating Node.js service:", serviceName)

	// Check if service already exists
	if _, err := os.Stat(serviceDir); !os.IsNotExist(err) {
		fmt.Println("❌ A service with this name already exists:", serviceName+"-service")
		return
	}

	// Create folder structor
	for _, folder := range folders {
		path := filepath.Join(serviceDir, folder)
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			fmt.Println("‼️ Failed to create:", path)
			return
		}
	}

	// Creating entry files
	CreateFile(filepath.Join(serviceDir, ".env"), "# .env file for "+serviceName)
	CreateFile(filepath.Join(serviceDir, "Dockerfile"), "# Docker file for "+serviceName)
	CreateFile(filepath.Join(serviceDir, "src/index.ts"), fmt.Sprintf("// %s-service entry\n", serviceName))

	fmt.Println("✅ Service created at:", serviceDir)
}
