/*
Copyright © 2023 Will Hutchinson will@thehutchery.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/whutchinson98/lambda-invoker-cli/cmd/invoke"
)

var version string = "0.0.0"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lambda-invoker-cli",
	Short: "Useful CLI for working with lambdas",
	Long:  `Useful CLI tool for working with lambdas to allow increased developer productivity`,
	Run: func(cmd *cobra.Command, args []string) {
		ver, _ := cmd.Flags().GetBool("version")
		if ver {
			fmt.Printf("Version: %s\n", version)
		}
	},
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
	cobra.OnInitialize()
	rootCmd.Flags().Bool("version", false, "Get the version of the CLI")
	rootCmd.AddCommand(invoke.InvokeCmd)
}
