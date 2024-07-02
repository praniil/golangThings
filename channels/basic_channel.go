package channels

import (
	"fmt"
	"sync"
	"time"
)

func Channels() {
	var message_chan = make(chan string)
	waitGrp := new(sync.WaitGroup)

	waitGrp.Add(1)
	go func(messageChan chan string) { 
		defer func() {
			println("this routine is done")
			waitGrp.Done()
		}()
		messageChan <- "ping" 
		println("before hello receiver not ready")
		messageChan <- "hello"
		println("after hello receiver ready")
	}(message_chan)

	any_var := <-message_chan
	time.Sleep(time.Second * 5)
	another_var := <-message_chan

	fmt.Println(any_var)
	fmt.Println(another_var)
	waitGrp.Wait()
}