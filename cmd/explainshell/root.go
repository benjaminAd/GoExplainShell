package explainshell

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "explainshell",
	Short: "A CLI to explain Shell commands",
	Long:  "A CLI to explain Shell commands using OpenAI's GPT-3 API",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. Something went wrong while executing the CLI: %v", err)
		os.Exit(1)
	}
}
