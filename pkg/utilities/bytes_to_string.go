package utilities

import (
	"fmt"
	"strings"
)

// Mac Mac converts a byte array to a mac address
func Mac(data []byte) string {
	return strings.Join(ConvertToHex(data), ":") 
}

// ConvertToHex Converts a byte array to its Hex representation
func ConvertToHex(data []byte) []string {	
	var hexBytesArray = make([]string, len(data))
	for i,b :=range data {
		hexBytesArray[i] = fmt.Sprintf("%X", b)
	}
	return hexBytesArray
}

// IP Converts a byte array to an IP address separated by decimal numbers and dots
func IP(data []byte) string {
	
	var ipPartAsString = make([]string, len(data))

	for i,b := range data {
		ipPartAsString[i] = fmt.Sprintf("%d", b)
	}
	
	return strings.Join(ipPartAsString, ".")
}