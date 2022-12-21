package cmd

import (
	"github.com/henriquegmendes/go-expert-client-server-api/server"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run server application",
	Run: func(cmd *cobra.Command, args []string) {
		server.InitServer()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
