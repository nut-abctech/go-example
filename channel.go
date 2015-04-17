package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; i < 10; i++ {
		c <- "ping"
	}
}

func print(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	//create a new channel with make (chan val-type)
	messages := make(chan string)
	go pinger(messages)
	go print(messages)
	//buffer channel
	buffmsg := make(chan string, 2)
	buffmsg <- "buffer"
	buffmsg <- "channel"

	fmt.Println(<-buffmsg)
	fmt.Println(<-buffmsg)

	var input string
	fmt.Scanln(&input)
}
