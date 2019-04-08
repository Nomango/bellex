// Copyright (C) 2018 Nomango - All Rights Reserved

package main

import (
	"fmt"
	"log"

	"github.com/nomango/bellex/server/modules/settings"
)

func main() {

	var serverIP string

	if settings.IsDevelopeMode() {
		serverIP = "127.0.0.1:7777"
	} else {
		serverIP = "132.232.126.221:7777"
	}

	client, err := NewClient(serverIP)
	if err != nil {
		fmt.Println("Start client failed: ", err)
		return
	}

	defer client.Close()

	go func() {
		receiver := client.Receiver()
		for {
			// handle response
			if response, err := receiver.ReadString(byte(0)); err != nil {
				log.Fatalln(err)
			} else {
				fmt.Println(response)
			}
		}
	}()

	fmt.Println("===============================")
	fmt.Println("Bellex Console v0.9")
	fmt.Println()

	fmt.Println("Menus:")
	fmt.Println("- 1: Send connect request")
	fmt.Println("- 0: Exit")

	fmt.Println()
	fmt.Println("Copyright (c) 2019 Bellex")
	fmt.Println("===============================")

	for {
		var cmd int
		if _, err := fmt.Scanln(&cmd); err != nil {
			fmt.Println("Unknown command")
			continue
		}

		switch cmd {
		case 0:
			return
		case 1:
			client.Send([]byte(`id:12345678;code:00000000;req:connect;`))
		default:
			fmt.Println("Unknown command")
		}
	}
}
