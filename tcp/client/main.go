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

	go client.Receive()

	// send request per second
	go func() {
		heartBeatTick := time.Tick(time.Second)
		for {
			select {
			case <-heartBeatTick:
				client.RequestServerTime()
			case <-client.Stop:
				return
			}
		}
	}()

	// wait for exit
	<-client.Stop
}
