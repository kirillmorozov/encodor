package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "Encodes your input",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ExecuteCustomIO(args []string, output io.Writer) error {
	rootCmd.SetArgs(args)
	rootCmd.SetOutput(output)
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
