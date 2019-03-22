// Copyright (C) 2018 Nomango - All Rights Reserved

package main

import (
	"fmt"

	"github.com/nomango/bellex/server/services/ntp"
)

func main() {
	guns := make(chan struct{}, len(ntp.Servers))
	results := make(chan struct{}, len(ntp.Servers))

	for _, host := range ntp.Servers {
		go func(host string) {
			// wait
			<-guns

			if now, err := ntp.SendRequest(host); err == nil {
				fmt.Println("Response from", host, now)
			}
			results <- struct{}{}
		}(host)
	}

	// Start
	for i := 0; i < len(ntp.Servers); i++ {
		guns <- struct{}{}
	}

	// Wait for responses
	for i := 0; i < len(ntp.Servers); i++ {
		<-results
	}
}
