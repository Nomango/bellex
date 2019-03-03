// Copyright (C) 2018 Nomango - All Rights Reserved

package main

import (
	"fmt"
	"time"
)

func main() {

	client, err := NewClient("132.232.126.221:7777")
	if err != nil {
		fmt.Println("Start client failed: ", err)
		return
	}

	defer client.Close()

	responseChan := make(chan struct{}, 3)
	heartBeatTick := time.Tick(time.Second)

	fmt.Println("Start to request server time")

	go func() {
		receiver := client.Receiver()
		for {
			// handle response
			if response, err := receiver.ReadString(byte(0)); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(response)
			}
		}
	}()

	// send request per second in goroutine
	for i := 0; i < cap(responseChan); i++ {
		select {
		case <-heartBeatTick:
			client.RequestTime()
			responseChan <- struct{}{}
		}
	}

	// wait response
	for i := 0; i < cap(responseChan); i++ {
		<-responseChan
	}
}
