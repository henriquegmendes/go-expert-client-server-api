package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "client-server-api",
	Short: "cli to run client-server-api application",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("error triggering application execution: %s", err.Error())
	}
}
