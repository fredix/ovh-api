package main

import (
	"fmt"
	"flag"
	"os"
	
	"github.com/ovh/go-ovh/ovh"
)

func main() {
	endPoint := flag.String("endpoint", "ovh-eu", "example ovh-eu")
	ip := flag.String("ip", "xx.xx.xx.xx", "your ip failover")

	flag.Parse()
	
	name, err := os.Hostname()
         if err != nil {
                 panic(err)
         }
         fmt.Println("Hostname reported by kernel : ", name)


	client, err := ovh.NewEndpointClient(*endPoint)
	if err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}

	type routedTo struct {
		RoutedTo string `json:"to"`
	}

	to := &routedTo{RoutedTo: name}

	if err := client.Post("/ip/" + *ip + "/move", to, nil); err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}

}
