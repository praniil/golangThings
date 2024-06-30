package main

import (
	"fmt"
	"learning_go/closures"
	"learning_go/goroutines"
)

func main() {
	fmt.Println("in main");
	// closures.IncCount();
	closures.Fibonacci();
	goroutines.AtomicCounter();
}