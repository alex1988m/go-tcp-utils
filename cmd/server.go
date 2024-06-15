/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github/alex1988m/go-tcp-utils/server"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start test server",
	RunE: func(cmd *cobra.Command, args []string) error {
		target := viper.GetString("target")
		server := &server.TestServer{Addr: target}
		if err := server.Start(); err != nil {
			return err	
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

}
