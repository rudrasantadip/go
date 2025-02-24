package main

import "fmt"

func main() {
    // Declare and initialize an anonymous struct.
    person := struct {
        Name string
        Age  int
    }{
        Name: "Alice",
        Age:  30,
    }

    fmt.Println(person)
}

