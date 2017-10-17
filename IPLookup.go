package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	host := flag.String("host", "", "host name")

	flag.Parse()

	if *host == "" {
		flag.Usage()
		os.Exit(1)
	} else {
		ips, err := net.LookupIP(*host)
		if err != nil {
			fmt.Printf("Could not get IPs: %v\n", err)
			os.Exit(1)
		}
		for _, ip := range ips {
			fmt.Printf(*host+" IN A %s\n", ip)
		}
	}
}
