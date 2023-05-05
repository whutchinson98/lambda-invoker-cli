package invoke

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/whutchinson98/lambda-invoker-cli/pkg/invoker"
)

var InvokeCmd = &cobra.Command{
	Use:   "invoke",
	Short: "Invokes a lambda to simulate an API call",
	Long: `Invokes a lambda with the provided json file to simulate an API call.
	This allows users to deploy lambdas and easily test them via this CLI.`,
	Args: cobra.ExactArgs(1), // Name of lambda to be invoked
	Run: func(cmd *cobra.Command, args []string) {
		lambdaName := args[0]

		region, _ := cmd.Flags().GetString("region")
		request, _ := cmd.Flags().GetString("request")
		workingDir, _ := os.Getwd()

		invoker.InvokeLambdaCmd(context.TODO(), &invoker.LambdaInvokerConfig{
			LambdaName: lambdaName,
			Region:     region,
			Request:    request,
			WorkingDir: workingDir,
		})
	},
}

func InitInvoke() {
	InvokeCmd.Flags().String("region", "us-east-1", "The AWS region the Secret is located")
	InvokeCmd.Flags().StringP("request", "r", "./request.json", "Relative path to your request file")
}
