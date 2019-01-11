package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "iaws",
	Short: "AWS Insights CLI",
	Long: `AWS Insights CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("root called")
	},
}

func Execute() {
	rootCmd.SetOutput(os.Stdout)
	if err := rootCmd.Execute(); err != nil {
		rootCmd.SetOutput(os.Stderr)
		rootCmd.Println(err)
		os.Exit(1)
	}
}