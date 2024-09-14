package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate code or files",
	Long:  `This command runs 'go generate ./...' to generate code or files based on directives in your Go source files.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Execute the 'go generate ./...' command
		err := runGoGenerate()
		if err != nil {
			fmt.Println("Error running go generate:", err)
		} else {
			fmt.Println("Code generation completed successfully.")
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

// runGoGenerate runs the 'go generate ./...' command
func runGoGenerate() error {
	// Create the command
	generateCmd := exec.Command("go", "generate", "./...")

	// Capture output (optional)
	output, err := generateCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%v: %s", err, string(output))
	}

	// Print the output
	fmt.Println(string(output))
	return nil
}
