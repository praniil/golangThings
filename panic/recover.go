package panic

import "fmt"

func anyPanicFn() {
	panic("panic occurred")
}

func PanicAndRecover() {
	defer func() {
		if r:= recover();
		r!= nil {
			fmt.Println("Recovered, error: \n", r)
		}
	}()
	anyPanicFn();
	//this doesnt print because panic occurs and main fn execution stops and only the deferred closure is active
	println("after panic");
}