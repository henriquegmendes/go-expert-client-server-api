package cmd

import (
	"github.com/henriquegmendes/go-expert-client-server-api/client"
	"github.com/spf13/cobra"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "run client application",
	Run: func(cmd *cobra.Command, args []string) {
		client.InitClient()
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
