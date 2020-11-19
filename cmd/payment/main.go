package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Started...")
	go func() {
		for {
			time.Sleep(time.Second*10)
			fmt.Println("waiting.")
		}
	}()
	wait := make(chan bool)
	<-wait
}
