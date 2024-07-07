package pipelining

import (
	"fmt"
	"time"
)

func getNumber(nums ...int) chan int {
	chan1 := make(chan int)
	go func () {
		for _, num:= range nums {
			chan1 <- num
		}
		close(chan1)
	}() 
	return chan1
}

func squareNumber(in chan int) chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			out <- num * num
		}
		close(out)
	} ()
	return out
}

func calculateSquare(nums ...int) {
	startTime := time.Now()
	for _, num := range nums{
		println("the square is: ", num * num)
	}
	stopTime := time.Now();
	fmt.Println("time taken without pipelining: ", stopTime.Sub(startTime))
}

func MainFunc () {
	calculateSquare(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13)
	startTime := time.Now()
	c := getNumber(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13);
	out := squareNumber(c)

	for num := range out {
		fmt.Printf("the square is: %d \n", num)
	}
	stopTime := time.Now();
	fmt.Println("time taken: ", stopTime.Sub(startTime))
}