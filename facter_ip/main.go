// A simple exe to get all the network interface addresses for facter
package main

import (
	"fmt"
	"net"
)

func getIfaceAddr() {
	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for addrIndex, address := range addrs {
			simpleAddr, _, _ := net.ParseCIDR(address.String())
			fmt.Printf("ip_addr_%s_%d=%s\n", i.Name, addrIndex, simpleAddr) //i.Name is a field of interfaces struct
		}
	}
}

func main() {
	getIfaceAddr()
}
