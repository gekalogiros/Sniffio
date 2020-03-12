package cmd

import (
	"io"
  "fmt"
  "github.com/spf13/cobra"
)


// NewVersionCommand creates the version command
func NewVersionCommand(out io.Writer) * cobra.Command {
	return &cobra.Command{
    Use:   "version",
    Short: "Print the version number of Sniffio",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("Sniffio - network traffic monitoring v0.1 -- HEAD")
    }, 
  }
}