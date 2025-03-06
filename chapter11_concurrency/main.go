package main

import (
	"fmt"
	"time"
)

/*
Often it happens, we need to run tasks in parallely or concurrently, for example
when multiple users try to access a resource, we create seperate instances for them such that
one user need not wait for the other... here we need concurrency
Go has a special feature called goroutines, which lets functions run concurrently.
It uses less memory than traditional threads.

*** INORDER TO CONVERT A FUNCTION INTO A GOROUTINE -> we just preceed it with "go" before calling the function
func display(){
...logic
}

display() // normal blocked approach
go display() // concurrently running
*/

type User struct {
	username string
	password string
}

func printUserDetails(user User) {
	fmt.Printf("username: %s\npassword: %s\n", user.username, user.password)
}

func main() {
	u1 := User{"rudrasantadip", "abc123"}
	u2 := User{"pokemonash", "pika123"}

	for {
		printUserDetails(u1)
		time.Sleep(500 * time.Millisecond)
		go printUserDetails(u2) // runs parallelly as a go routine

		/*
			in absence of the sleep function, the go routine will not be printed since it
			will run parallery and the program will not wait for it to complete,
			so we initiate a delay to syncronize the two function calls.
		*/
	}

}
