package createService

import (
	"fmt"
	"os"
)

// createFile creates a file at the given path and writes the content to it.
func CreateFile(path string, content string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("❌ Could not create file:", path)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("❌ Could not write to file:", path)
	}
}
