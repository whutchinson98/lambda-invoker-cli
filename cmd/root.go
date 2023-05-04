/*
Copyright © 2023 Will Hutchinson will@thehutchery.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lambda-invoker-cli",
	Short: "Invokes a lambda to simulate an API call",
	Long: `Invokes a lambda with the provided json file to simulate an API call.
	This allows users to deploy lambdas and easily test them via this CLI.`,
	Run: func(cmd *cobra.Command, args []string) {},
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
	// TODO: Add this!!
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.lambda-invoker-cli.yaml)")
	rootCmd.Flags().String("region", "us-east-1", "The AWS region the Secret is located")
	rootCmd.Flags().StringP("request", "r", "./request.json", "Relative path to your request file")
}
