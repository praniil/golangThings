package main

import (
	"fmt"
	// "learning_go/closures"
	"learning_go/goroutines"
	"learning_go/panic"
)

func main() {
	fmt.Println("in main");
	// closures.IncCount();
	// closures.Fibonacci();
	// goroutines.AtomicCounter();
	goroutines.Mutex();
	panic.PanicAndRecover();

}