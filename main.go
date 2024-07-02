package main

import (
	"bufpool/buff"
	"fmt"
)

func main() {

	// creates a buffer pool
	buffPool := buff.NewBufferPool(2000)

	// retrieve the buffer from the pool
	buf := buffPool.Get()

	// wrtie data to the pool
	buf.Write([]byte("hello,world!"))
	fmt.Println(buf)

	// after writing the data successfully, return the buffer to the pool
	buffPool.Put(buf)

}
