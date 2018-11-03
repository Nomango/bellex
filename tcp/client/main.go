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

	go func() {
		responseChan := make(chan struct{}, 5)
		heartBeatTick := time.Tick(time.Second)

		fmt.Println("Start to request server time")

		// send request per second in goroutine
		for i := 0; i < cap(responseChan); i++ {
			select {
			case <-heartBeatTick:
				client.RequestTime()
				// handle response
				go func() {
					fmt.Print("Server response: ", client.Receive())
					responseChan <- struct{}{}
				}()
			}
		}

		// wait response
		for i := 0; i < cap(responseChan); i++ {
			<-responseChan
		}
		client.Close()
	}()

	// wait until tcp connection is closed
	client.WaitForClosed()
}
