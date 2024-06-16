package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:          "go-tcp-utils",
	Short:        "a set of useful tcp/http utils",
	SilenceUsage: true,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() {
		viper.AddConfigPath(".")
		viper.SetConfigName(".cobra")
		viper.AutomaticEnv()
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	})

	rootCmd.PersistentFlags().StringP("target", "t", "localhost:5555", "target host:port")
	viper.BindPFlag("target", rootCmd.PersistentFlags().Lookup("target"))
	rootCmd.PersistentFlags().StringP("proxy", "p", "localhost:4444", "proxy host:port")
	viper.BindPFlag("proxy", rootCmd.PersistentFlags().Lookup("proxy"))
}
