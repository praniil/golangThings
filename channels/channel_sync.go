package channels

import (
	"fmt"
	"time"
)

func worker(done chan string) {
	time.Sleep(10* time.Millisecond)
	done <- "done"
}

func ChannelSync() {
	any_chan := make(chan string, 1)
	go worker(any_chan)
	any_string := <-any_chan
	fmt.Println(any_string);
}