package main

import (
	"fmt"
)

func main() {
	// The way for synchronizing and communicating different go routines
	// is using Channels.
	// Channels always have a type (any type)
	// Data in channel by given channel is the same type
	// The second element in channel is the capacity of channel
	c := make( chan string, 3)

	c <- "1"
	fmt.Println("1")

	c <- "2"
	fmt.Println("2")

	c <- "3"
	fmt.Println("3")

	c <- "4"
	fmt.Println("4")
	//go func(input chan string ) {
	//	fmt.Println("sending 1 to the channel")
	//	input <- "Hello1"
	//
	//	fmt.Println("sending 2 to the channel")
	//	input <- "Hello2"
	//
	//	fmt.Println("sending 3 to the channel")
	//	input <- "Hello3"
	//
	//	fmt.Println("sending 4 to the channel")
	//	input <- "Hello4"
	//}(c)
	//
	//fmt.Println("receiving from channel")
	////greeting := <- c
	//for greeting := range c {
	//	fmt.Println("greeting received")
	//	fmt.Println(greeting)
	//}




	// To run go routine you must put `go` in front of func you want concurrency
	// go HelloWorld()

}

func HelloWorld() {
	fmt.Println("Hello World")
}
