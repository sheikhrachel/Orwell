package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var spelunkerCmd = &cobra.Command{
	Use:   "spelunker",
	Short: "Test generation",
	Long: `Spelunker rapidly parses through an existing codebase and generates a test directory (if missing) and necessary test files and modules
Reference the flags below for the necessary options

Supported languages: python
Supported libraries: requests`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("spelunker called")
	},
}

func init() {
	rootCmd.AddCommand(spelunkerCmd)
}
