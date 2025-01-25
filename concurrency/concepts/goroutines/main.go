package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello("Hello!")
	sayHello("Hello, world!")
}

func sayHello(msg string) {
	for {
		fmt.Println(msg)
		time.Sleep(time.Second)
	}
}
