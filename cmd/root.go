package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func NewRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Short: "Encodes your input",
	}
	beghiloszCmd := NewBeghiloszCmd()
	rootCmd.AddCommand(beghiloszCmd)
	zalgoCmd := NewZalgoCmd()
	rootCmd.AddCommand(zalgoCmd)
	return rootCmd
}

func Execute() {
	rootCmd := NewRoot()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ExecuteCustomIO(args []string, output io.Writer) error {
	rootCmd := NewRoot()
	rootCmd.SetArgs(args)
	rootCmd.SetOutput(output)
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
