package cmd

import (
	"github.com/spf13/cobra"
)

// NewSniffioCommand Returns the main cli command
func NewSniffioCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "sniffio",
		Short: "Sniffio is a dead simple cli tool that helps you sniff network traffic",
		Long: `Sniffio is a cli tool that makes network sniffing dead simple.
The tool has been designed with simplicity in mind.`,
	}
}