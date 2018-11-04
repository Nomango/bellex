// Copyright (C) 2018 Nomango - All Rights Reserved

package main

import (
	"github.com/nomango/bellex/services/ntp"
)

func main() {
	requests := make(chan struct{}, 3)

	for _, host := range ntp.Servers {
		go func(host string) {
			ntp.SendRequest(host)
			requests <- struct{}{}
		}(host)
	}

	// Wait for exit
	for range make([]int, 3) {
		<-requests
	}
}
