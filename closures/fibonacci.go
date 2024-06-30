package closures

import "fmt"

func GiveFibo() func() int {
	firstItem := 0;
	secondItem := 0;
	return func () int {
		if firstItem == 0 {
			firstItem = 1;
			secondItem = 1;
			return 0;
		} else {
			current := firstItem;
			firstC := secondItem;
			secondItem = firstItem + secondItem;
			firstItem = firstC;
			return current
		}

	}
}

func Fibonacci() {
	fmt.Println("in fibo");
	for i:= 0; i < 10; i++ {
		fiboSeries := GiveFibo();
		fmt.Println(fiboSeries());
	}
}