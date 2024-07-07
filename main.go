package main

import (
	"fmt"
	// "learning_go/closures"
	// "learning_go/goroutines"
	// "learning_go/panic"
	// "learning_go/channels"
	// "learning_go/files"
	"learning_go/pipelining"
)

func main() {
	fmt.Println("in main");
	// closures.IncCount();
	// closures.Fibonacci();

	// goroutines.AtomicCounter();
	// goroutines.Mutex();

	// panic.PanicAndRecover();

	// channels.Channels();
	// channels.BufferedChan();
	// channels.ChannelSync();

	// files.ReadFiles();
	pipelining.MainFunc();
}