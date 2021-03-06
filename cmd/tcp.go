package cmd

import (
	"strconv"
	"io"
	// "net"
	"strings"
	"github.com/google/gopacket/layers"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/spf13/cobra"
)

// NewTCPCommand creates the tcp command
func NewTCPCommand(out io.Writer) * cobra.Command {
	var port int
	tcpCommand := &cobra.Command{
		Use:   "tcp",
		Short: "Sniffs http traffic at your local interface",
		Run: func(cmd *cobra.Command, args []string) {
			executePcap(out, port)
		},
	}

	tcpCommand.Flags().IntVarP(&port, "port", "p", 8080, "port to sniff traffic on")

	return tcpCommand
}

func executePcap(out io.Writer, port int){

	handle, err := pcap.OpenLive("lo0", 65536, true, pcap.BlockForever)

	if err != nil {
		panic(err)
	}

	defer handle.Close()
	
	// var filter string = "port 80 or port 443"
	// err = handle.SetBPFFilter(filter)
	// if err != nil {
	// 	panic(err)
	// }

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {

		if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {

			// for _, layer := range packet.Layers() {

			// 	fmt.Println("Layer", layer.LayerType())

			// 	layerType := layer.LayerType()

			// 	if layerType == packet.Layer(layers.LayerTypeIPv4).LayerType() {
					
			// 		ipv4Layer := packet.Layer(layers.LayerTypeIPv4)

			// 		ipv4, _ := ipv4Layer.(*layers.IPv4)

			// 		fmt.Printf("Source IP: %d \n", ipv4.SrcIP)
			// 		fmt.Printf("Destination IP: %d \n", ipv4.DstIP)
			// 	}
			// }

			tcp, _ := tcpLayer.(*layers.TCP)

			if fmt.Sprintf("%d", tcp.DstPort) == strconv.Itoa(port) {
				
				payload := string(tcp.Payload);

				if strings.Contains(payload, "HTTP") {
					fmt.Printf("Source Port: %d \n", tcp.SrcPort)
					fmt.Printf("Destination Port: %d \n", tcp.DstPort)
					fmt.Printf("Acknowledgment: %d \n", tcp.Ack)
					fmt.Printf("Data Offset: %d \n", tcp.DataOffset)
					fmt.Printf("Packet Contents: %d \n", tcp.Contents)
					fmt.Printf("Packet Payload: %d \n", tcp.LayerPayload())
					fmt.Printf("Packet Payload: \n %s \n", payload)
					fmt.Printf("Is SYN: %t \n", tcp.SYN)
					fmt.Printf("IS ACK: %t \n", tcp.ACK)
					fmt.Printf("IS FIN: %t \n", tcp.FIN)
				}

				// if tcp.SYN {
				// 	fmt.Println("[SYN] ")
				// }
				// if tcp.FIN {
				// 	fmt.Println("[FIN] ")
				// }
			}
		}
	}
}


  
  