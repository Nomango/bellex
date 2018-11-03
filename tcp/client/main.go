// Copyright (C) 2018 Nomango - All Rights Reserved

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nomango/bellex/services/tcp"
)

func main() {

	client, err := tcp.NewClient()
	if err != nil {
		fmt.Println("Start client failed: ", err)
		os.Exit(1)
	}

	// send request 5 times per second in goroutine
	go func() {
		heartBeatTick := time.Tick(time.Second)

		for i := range make([]int, 5) {
			select {
			case <-heartBeatTick:
				client.RequestServerTime()

				// handle response
				go func(i int) {
					response := client.Receive()
					fmt.Print("Server response: ", response)
					if i == 4 {
						client.Close()
					}
				}(i)
			}
		}
	}()

	// wait until tcp connection is closed
	client.WaitForClosed()
}
