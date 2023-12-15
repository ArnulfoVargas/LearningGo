package channels

import (
	"fmt"
	"time"
)

func ChannelsTest() {
	//! Will throw error

	//Channel declaration
	var channel chan string// Works by reference, poiting to nil
	
	go func() {
		channel <- "Some data"
	}() 

	// Close a channel
	// ? close(channel)

	//  Allows to share strings
	channel = make(chan string) // it can also be declared with make
	// Assign values
	
	go func() {
		channel <- "Hello world"
	}()
	// read values
	data := <-channel

	fmt.Println(data)

	// Usage

	go func ()  {
		channel <- "Bonjour"	
	}()

	data = <- channel

	fmt.Println(data)

	// Read only channels
	// ? var channel2 <-chan string
	// Write only channels
	// ? var channel3 chan<- string

	// usage 
	main2()
}
func sender(ch chan<- int){
	for i := 0; i < 3; i++ {
		ch <- i 
		fmt.Println("Sended:", i)
	}
}

func reciver(ch <-chan int){
	for i := 0; i < 3; i++ {
		fmt.Println("Received:", <-ch)
	}
}
func main2(){
	channel := make(chan int)

	go sender(channel)
	reciver(channel)

	// Channels with buffer
	main3()
}

// usage

func sender2(ch chan<- int){
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Println("Sended:", i)
	}
	close(ch)
}

// Looping through data with a for loop
func receiver2(ch <-chan int){
	for data := range ch{
		time.Sleep(time.Second * 1)
		fmt.Println("Received:",data)
	}
}

func main3(){
	// Channel with buffer declaration
	channel := make(chan int, 5)

	go sender2(channel)
	receiver2(channel)
	main4()
}

// design pattern to wait for asyncronous funcs
func asyncFunc() <-chan struct{} {
	ch := make(chan struct{})

	go func() {
		fmt.Println("Doing things...")	
		for i := 0; i < 3; i++ {
			fmt.Println(i, "...")
			time.Sleep(time.Second * 1)
		}
		fmt.Println("End")
		close(ch)
	}()

	return ch
}

func main4(){
	wait := asyncFunc()

	<-wait

	fmt.Println("Main 4 end")
}