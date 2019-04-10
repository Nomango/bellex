// Copyright (C) 2018 Nomango - All Rights Reserved

package main

import (
	"fmt"
	"log"
	"strings"
)

var (
	MechineCode   = "12345678"
	MechineSecret = "00000000"
)

func makeRequest(request string, data string) []byte {
	requestStr := "id:" + MechineCode + ";code:" + MechineSecret + ";req:" + request + ";data:" + data + ";"
	fmt.Println("发送数据包", requestStr)
	return []byte(requestStr)
}

func main() {

	var serverIP string

	fmt.Println("连接到本地还是远程TCP服务器? (0 or 1)")
	for {
		var cmd int
		if _, err := fmt.Scanln(&cmd); err != nil {
			continue
		}
		if cmd == 0 {
			serverIP = "127.0.0.1:7777"
		} else {
			serverIP = "132.232.126.221:7777"
		}
		break
	}

	client, err := NewClient(serverIP)
	if err != nil {
		fmt.Println("连接TCP服务器失败: ", err)
		return
	}

	defer client.Close()

	go func() {
		receiver := client.Receiver()
		for {
			// handle response
			response, err := receiver.ReadString(byte(0))
			if err != nil {
				log.Fatalln("TCP连接已断开", err)
				return
			}
			fmt.Printf("接收到数据包 (字节: %d) %s\n", len(response)-1, response)

			switch {
			case strings.Contains(response, "unique_code:") && len(response) == 22:
				MechineSecret = response[12:20]
				fmt.Println("更新主控机密码", MechineSecret)
			}
		}
	}()

	fmt.Println("===============================")
	fmt.Println("Bellex Console v0.9")
	fmt.Println()

	fmt.Println("Menus:")
	fmt.Println("- 1: 发送连接请求")
	fmt.Println("- 2: 发送校时请求")
	fmt.Println("- 3: 发送心跳包（工作状态）")
	fmt.Println("- 4: 发送心跳包（待机状态）")
	fmt.Println("- 5: 发送获取时间表请求")
	fmt.Println("- 0: 退出")

	fmt.Println()
	fmt.Println("Copyright (c) 2019 Bellex")
	fmt.Println("===============================")

	for {
		var cmd int
		if _, err := fmt.Scanln(&cmd); err != nil {
			fmt.Println("未知命令，请重新输入")
			continue
		}

		switch cmd {
		case 0:
			return
		case 1:
			client.Send(makeRequest("connect", ""))
		case 2:
			client.Send(makeRequest("proof_time", ""))
		case 3:
			client.Send(makeRequest("heart_beat", "status:working"))
		case 4:
			client.Send(makeRequest("heart_beat", "status:idle"))
		case 5:
			client.Send(makeRequest("get_schedule", ""))
		default:
			fmt.Println("未知命令，请重新输入")
		}
	}
}
