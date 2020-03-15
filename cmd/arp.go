package cmd

import (
	"github.com/gekalogiros/sniffio/pkg/arp"
	"io"
	"github.com/spf13/cobra"
)

// NewArpCommand creates the arp command
func NewArpCommand(out io.Writer) * cobra.Command {

	arpCommand := &cobra.Command{
		Use:   "arp",
		Short: "Sniffs arp traffic",
	}

	traffic := &cobra.Command{
		Use:   "traffic",
		Short: "Request/Reply arp packets",
		Run: func(cmd *cobra.Command, args []string) {
			arp.Sniff(out)
		},
	}

	devices := &cobra.Command{
		Use:   "devices",
		Short: "Finds devices in your local network",
		Run: func(cmd *cobra.Command, args []string) {
			arp.Detect(out)
		},
	}

	arpCommand.AddCommand(
		traffic, 
		devices,
	)

	return arpCommand
}



