package createService

import (
	"fmt"

	"github.com/spf13/cobra"
)

var lang string

var CreateServiceCmd = &cobra.Command{
	Use:   "create-service [name]",
	Short: "Scaffold a new microservice (Node.js or Go)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//serviceName := args[0]

		switch lang {
		case "node":
			// createNodeService(serviceName)
		case "go":
			// createGoService(serviceName)
		default:
			fmt.Println("❌ Unsupported language:", lang)
		}
	},
}

func init() {
	CreateServiceCmd.Flags().StringVarP(&lang, "lang", "1", "node", "Programming language: node | go")
}
