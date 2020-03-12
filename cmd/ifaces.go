package cmd

import (
	"fmt"
	"io"
  	"github.com/google/gopacket/pcap"
  	"github.com/spf13/cobra"
)

// NewIfacesCommand creates the ifaces sub-command
func NewIfacesCommand(out io.Writer) * cobra.Command {
	return &cobra.Command{
		Use:   "interface",
		Short: "List all Network Interfaces that exist in the current machine",
		Run: func(cmd *cobra.Command, args []string) {
			findNetworkInterfaces(out)
		},
	}
}

func findNetworkInterfaces(out io.Writer) {
	
	interfaces, err := pcap.FindAllDevs()
	  
	if err != nil {
		panic(err)
	}

	for _, iface := range interfaces {
		fmt.Fprintf(out, "%s\n", iface.Name)
	}
}