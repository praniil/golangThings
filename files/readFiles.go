package files

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFiles() {
	data, err := os.ReadFile("/home/pranil/goProjects/learning_go/anyFile.txt")
	check(err)

	fmt.Println(string(data))

	//for more control over the file open the file in open mode

	file, err := os.Open("/home/pranil/goProjects/learning_go/anyFile.txt")
	check(err)

	//read 5 byte from beginning
	b1 := make([]byte, 5)
	n1, err := file.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s \n", n1, string(b1[:n1]))

	//using seek
	o2, err := file.Seek(6, io.SeekStart)
	check(err)
	b2 := make([]byte, 2)
	n2, err := file.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := file.Seek(1, io.SeekStart)
	check(err)
	b3:= make([]byte, 2)
	n3, err := io.ReadAtLeast(file, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	r4 := bufio.NewReader(file)
	b4, err := r4.Peek(7)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
	file.Close()

}