package cmd

import (
	"io"
	"net"
	"bytes"
	"strings"
	"time"
	"fmt"
	"encoding/binary"
	"strconv"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/spf13/cobra"
	. "github.com/logrusorgru/aurora"
)

// NewArpCommand creates the arp command
func NewArpCommand(out io.Writer) * cobra.Command {
	return &cobra.Command{
		Use:   "arp",
		Short: "Sniff arp traffic",
		Run: func(cmd *cobra.Command, args []string) {
			sniff(out)
		},
	}
}

func sniff(out io.Writer){

	handle, err := pcap.OpenLive("en0", 65536, false, pcap.BlockForever)

	if err != nil {
		panic(err)
	}

	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {

		arpLayer := packet.Layer(layers.LayerTypeARP)

		if arpLayer != nil {

			arp, _ := arpLayer.(*layers.ARP)

			s := `Operation %s
Source: %s - %s
Destination: %s - %s

`
			operationType := arpOperation(arp.Operation)

			var operationTypeColoured interface {} 

			if operationType == "Request" {
				operationTypeColoured = Green(operationType)
			} else {
				operationTypeColoured = Red(operationType)
			}

			fmt.Fprintf(out, s, operationTypeColoured,
				Green(ip(arp.SourceProtAddress)),
				Green(mac(arp.SourceHwAddress)),
				Green(ip(arp.DstProtAddress)),
				Green(mac(arp.DstHwAddress)),
			)
		}
	}
}

func mac(data []byte) string {
	return strings.Join(convertToHex(data), ":") 
}

func ip(data []byte) string {
	
	var ipPartAsString = make([]string, len(data))

	for i,b := range data {
		ipPartAsString[i] = fmt.Sprintf("%d", b)
	}
	
	return strings.Join(ipPartAsString, ".")
}

func arpOperation(operation uint16) string {
	operationAsString := ""
	if operation == 1 {
		operationAsString = "Request"
	} else if operation == 2 {
		operationAsString = "Reply"
	} else {
		panic("ERROR-1")
	}
	
	return operationAsString
}

func detect(){
	ifaces, err := net.Interfaces()

	if (err != nil){
		panic(err)
	}

	if err == nil {

		for _, iface := range ifaces {

			var addrs, err = iface.Addrs();

			if (err != nil) {
				panic(err)
			}
			
			for _, addr := range addrs {
				
				ip, ipnet, err := net.ParseCIDR(addr.String())

				if err != nil {
					panic(err)
				}

				ipv4 := ip.To4()
				mask := ipnet.Mask[len(ipnet.Mask)-4:]

				if (ipv4 == nil || ipv4[0] == 127 || (mask[0] != 0xff || mask[1] != 0xff)){
					continue
				}

				println("Interface: " + iface.Name + ", Network: " + ipnet.String() + ", IP: " + ipv4.String())

				handle, err := pcap.OpenLive(iface.Name, 65536, true, pcap.BlockForever)
				if err != nil {
					panic(err)
				}
				defer handle.Close()

				stop := make(chan struct{})
				
				go read(handle, &iface, stop)
				defer close(stop)

				for {

					if err := write(handle, iface, *ipnet); err != nil {
						println("error writing packets on %v: %v", iface.Name, err)
					}

					time.Sleep(10 * time.Second)
				}
			}
		}
	} 
}

func read(handle *pcap.Handle, iface *net.Interface, stop chan struct{}) {
	in := gopacket.NewPacketSource(handle, layers.LayerTypeEthernet).Packets()
	for {
		var packet gopacket.Packet
		select {
		case <- stop:
			return
		case packet = <- in:
			arpLayer := packet.Layer(layers.LayerTypeARP)
			if arpLayer == nil {
				continue
			}
			arp := arpLayer.(*layers.ARP)
			if arp.Operation != layers.ARPReply || bytes.Equal([]byte(iface.HardwareAddr), arp.SourceHwAddress) {
				continue
			}
			println("IP: " + net.IP(arp.SourceProtAddress).String() + ", Device: " + net.HardwareAddr(arp.SourceHwAddress).String())
		}
	}
}

func write(handle *pcap.Handle, iface net.Interface, addr net.IPNet) error {
	eth := layers.Ethernet {
		SrcMAC:       iface.HardwareAddr,
		DstMAC:       net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		EthernetType: layers.EthernetTypeARP,
	}
	arp := layers.ARP {
		AddrType:          layers.LinkTypeEthernet,
		Protocol:          layers.EthernetTypeIPv4,
		HwAddressSize:     6,
		ProtAddressSize:   4,
		Operation:         layers.ARPRequest,
		SourceHwAddress:   []byte(iface.HardwareAddr),
		SourceProtAddress: []byte(addr.IP),
		DstHwAddress:      []byte{0, 0, 0, 0, 0, 0},
	}
	buffer := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}
	
	gopacket.SerializeLayers(buffer, opts, &eth, &arp)

	generatedIps := ips(&addr)

	size := len(generatedIps)
	
	println("size" + strconv.Itoa(size))

	for _, ip := range generatedIps {

		fmt.Println(ip)

		// println("Write to" + ip.String())

		// arp.DstProtAddress = []byte(ip)
		
		// gopacket.SerializeLayers(buffer, opts, &eth, &arp)

		// if err := handle.WritePacketData(buffer.Bytes()); err != nil {
		// 	return err
		// }
	}
	
	return nil
}

func ips(n *net.IPNet) (out []net.IP) {
	num := binary.BigEndian.Uint32([]byte(n.IP))
	mask := binary.BigEndian.Uint32([]byte(n.Mask))
	num &= mask
	for mask < 0xffffffff {
		var buf [4]byte
		binary.BigEndian.PutUint32(buf[:], num)
		out = append(out, net.IP(buf[:]))
		mask++
		num++
	}
	return
}

func increment(ip net.IP) {
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		if ip[i] != 0 {
			break
		}
	}
}
  
  