/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github/alex1988m/go-tcp-utils/server"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// testHttpCmd represents the testHttp command
var testHttpCmd = &cobra.Command{
	Use:   "test-http",
	Short: "create test http server",
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
	serverCmd.AddCommand(testHttpCmd)
}
