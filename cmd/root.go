package cmd

import (
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
		rootCmd.PrintErr(err)
		os.Exit(1)
	}
}
