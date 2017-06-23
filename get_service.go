package main

import (
	"fmt"
	"flag"
	
	"github.com/ovh/go-ovh/ovh"
)

func main() {

	endPoint := flag.String("endpoint", "ovh-eu", "example ovh-eu")
	ip := flag.String("ip", "xx.xx.xx.xx", "your ip failover")

	flag.Parse()
	
	client, err := ovh.NewEndpointClient(*endPoint)
	if err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}

	// Get all the ip services
	//ipInfo := []string{}
	type routedTo struct {
		ServiceName string `json:"serviceName"`
	}
	type ipInfo struct {
		Country string `json:"country"`
		RoutedTo routedTo
	}

        sn := ipInfo{}

	if err := client.Get("/ip/" + *ip, &sn); err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}


	fmt.Println("serviceName : ", sn)
}
