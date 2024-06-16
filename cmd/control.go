/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github/alex1988m/go-tcp-utils/handler"
	"github/alex1988m/go-tcp-utils/server"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// controlCmd represents the control command
var controlCmd = &cobra.Command{
	Use:   "control",
	Short: "start proxy tcp server which expose bash for client",
	RunE: func(cmd *cobra.Command, args []string) error {
		proxyAddr := viper.GetString("proxy")
		var handler handler.TCPHandler = &handler.TCPControlHandler{Command: []string{"/bin/sh", "-i"}}
		server := &server.TCPServer{
			ServerAddr: proxyAddr,
			Handler:    handler,
		}

		if err := server.Start(); err != nil {
			return err
		}
		defer server.Stop()
		if err := server.Serve(); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	serverCmd.AddCommand(controlCmd)
}
