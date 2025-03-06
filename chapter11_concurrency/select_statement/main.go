package main

import (
	"fmt"
	"time"
)

/*
in real life scenarios we can be recieving data from multiple locations, and at different
points of time like both the users are logged in but someone is sending the data while the
other is not, here we use the select statement which waits on multiple "channels" and it
allows us to proceed with which ever channel is ready making it ideal for handling asynchronous tasks efficiently.
*/

type User struct {
	username string
	password string
}

func sendDataU1(user User, ch chan string) {
	time.Sleep(2 * time.Second) //pauses for 2 seconds
	ch <- user.username + " is sending data"
}

func sendDataU2(user User, ch chan string) {
	time.Sleep(4 * time.Second) //pauses for 4 seconds
	ch <- user.username + " is sending data"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go sendDataU1(User{"user1", "password123"}, ch1)
	go sendDataU2(User{"user2", "password123"}, ch2)

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1) // this one is finised faster cause channel 1 is finished within 2 seconds
	case msg2 := <-ch2:
		fmt.Println(msg2)
	/*
		The default is used is case no channel responds,
		in case the select may block the program hence default
		is executed to terminate it
	*/
	default:
		fmt.Println("No channels are ready")
	}

	/*
		The output will be "user1 sends data",
		because it is being finished faster

		select waits until at least one channel operation is ready.
		If multiple cases are ready, one is chosen at random.
		The default case executes if no other case is ready, avoiding a block.
	*/
}
