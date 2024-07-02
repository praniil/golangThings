package channels

func BufferedChan() {
	bufferedChan := make(chan string, 2)
	bufferedChan <- "string1"
	bufferedChan <- "string2"
	// println(<-bufferedChan)
	// we can only add something in the buffer only if it has empty space, otherwise wait for the chan to send the data, create the space then send the data to the buffer
	println(<-bufferedChan)
	bufferedChan <- "extra"
	println(<-bufferedChan)
	println(<-bufferedChan)
}
