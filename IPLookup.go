package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	revFlag := flag.Bool("r", false, "reverse lookup")

	flag.Parse()
	param := flag.Args()

	if len(param) > 1 {
		fmt.Printf("You must specify only 1 command line parameter")
	} else {
		if *revFlag {

			addresses, err := net.LookupAddr(param[0])
			if err != nil {
				fmt.Printf("Could not get Addresses: %v\n", err)
				os.Exit(1)
			}
			for _, address := range addresses {
				fmt.Printf(param[0]+" IN A %s\n", address)
			}
		} else {
			ips, err := net.LookupIP(param[0])
			if err != nil {
				fmt.Printf("Could not get IPs: %v\n", err)
				os.Exit(1)
			}
			for _, ip := range ips {
				fmt.Printf(param[0]+" IN A %s\n", ip)
			}
		}
	}
}
