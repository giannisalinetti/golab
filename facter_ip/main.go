// A simple exe to get all the network interface addresses for Facter.
// By default Facter provides only the first ip address of every interface on the machine.
// This program adds and extra output in the format 'ip_addr_<iface>_N' where N is the
// N-th ip address associated.

// Build and install in /etc/facter/facts.d
  
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
			fmt.Printf("ip_addr_%s_%d=%s\n", i.Name, addrIndex, simpleAddr) 
		}
	}
}

func main() {
	getIfaceAddr()
}
