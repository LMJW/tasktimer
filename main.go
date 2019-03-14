package main

import (
	"fmt"
	"time"
)

func main() {
	t := NewTask("a", "timer", 10*time.Second, time.Second)
	go t.Start()
	time.Sleep(5 * time.Second)
	t.Pause(true)
	time.Sleep(5 * time.Second)
	t.Pause(false)
	s := make(chan bool)
	// time.Sleep(10 * time.Second)
	t.Extend(10 * time.Second)

	select {
	case <-s:
		fmt.Println("?")
	}
}
