package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string, 1)

	receive(c, "foo")
	fmt.Println(<-c)

	var message string
	c <- "bar"
	send(c, &message)
	fmt.Println(message)

	var cr <-chan string
	cr = receivableChan("ok")
	fmt.Println(<-cr)

	var cs chan<- string
	cs = sendableChan()
	cs <- "one"
	time.Sleep(100 * time.Millisecond)
	cs <- "two"
	time.Sleep(100 * time.Millisecond)
	cs <- "three"
	time.Sleep(100 * time.Millisecond)
}

// チャンネル "が" 受け取る
// チャンネル "に" 送る
func receive(c chan<- string, message string) {
	c <- message
}

// チャンネル "から" 送る
// チャンネル "から" 受け取る
func send(c <-chan string, message *string) {
	*message = <-c
}

// 受信専用チャンネルの返却
func receivableChan(message string) <-chan string {
	c := make(chan string, 1)
	c <- message
	return c
}

// 送信専用チャンネルの返却
func sendableChan() chan<- string {
	c := make(chan string, 1)
	go func(cr <-chan string) {
		for message := range cr {
			fmt.Println(message)
		}
	}(c)
	return c
}
