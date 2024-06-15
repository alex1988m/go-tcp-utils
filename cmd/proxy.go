package cmd

import (
	"github/alex1988m/go-tcp-utils/server"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// proxyCmd represents the proxy command
var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "creates proxy server for the given tcp server",
	RunE: func(cmd *cobra.Command, args []string) error {
		proxyAddr := viper.GetString("proxy")
		targetAddr := viper.GetString("target")

		log.Printf("proxy addr: %s, target addr: %s", proxyAddr, targetAddr)

		proxy := &server.ProxyServer{
			ProxyAddr:  proxyAddr,
			TargetAddr: targetAddr,
		}

		if err := proxy.Start(); err != nil {
			return err
		}
		defer proxy.Stop()
		if err := proxy.Serve(); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(proxyCmd)

	proxyCmd.PersistentFlags().StringP("proxy", "p", "localhost:4444", "proxy host:port")

	viper.BindPFlag("proxy", proxyCmd.PersistentFlags().Lookup("proxy"))
}
