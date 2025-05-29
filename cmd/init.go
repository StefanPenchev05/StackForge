package cmd

import (
	"github.com/StefanPenchev05/StackForge/internal/scaffolder"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init <project-name>",
	Short: "Initializes a new full-stack microservice project",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		opts := scaffolder.InitOptions{
			ProjectName:  args[0],
			WithFrontend: withFrontend,
			WithAuth:     withAuth,
			Lang:         lang,
			WithDocker:   withDocker,
			WithCI:       withCI,
		}

		return scaffolder.InitProject(opts)
	},
}

var (
	withFrontend bool
	withAuth     bool
	lang         string
	withDocker   bool
	withCI       bool
)

func init() {
	initCmd.Flags().BoolVar(&withFrontend, "with-frontend", false, "--with-frontend=true")
	initCmd.Flags().BoolVar(&withAuth, "with-auth", false, "--with-auth=true")
	initCmd.Flags().StringVar(&lang, "lang", "node", "Backend language: node or go")
	initCmd.Flags().BoolVar(&withDocker, "with-docker", false, "--with-dcoker=true")
	initCmd.Flags().BoolVar(&withCI, "with-ci", false, "--with-ci=true")
	rootCmd.AddCommand(initCmd)
}
