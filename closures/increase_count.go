package closures

import "fmt"

func anyFunction() func() int{
	i:= 0
	return func() int {
		i++
		return i
	}
}

func IncCount () {
	fmt.Println("in inc counter");
	nextInt := anyFunction();
	fmt.Println(nextInt());
	fmt.Println(nextInt());

	anotherInt := anyFunction();
	fmt.Println(anotherInt());
	fmt.Println(anotherInt());
}