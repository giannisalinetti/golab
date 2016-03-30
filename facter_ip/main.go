// A simple exe to get all the network interface addresses for facter
// TODO: strip ip address from CIDR notation
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
			// address, subnet := net.ParseCIDR(cidrRes)
			fmt.Printf("ip_addr_%s_%d=%s\n", i.Name, addrIndex, address)
		}
	}
}

func main() {
	getIfaceAddr()
}
