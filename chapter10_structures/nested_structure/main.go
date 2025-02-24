package main

import "fmt"

/*
In a nested structure we can define a struct inside another struct
*/


//here we defined a user as a type
type User struct {
username string
password string
}


// in the product struct we are using the user as another attribute
type Product struct {
productId uint
users [] User
}

// initializing the struct
func main(){
	p := Product{
	1,
	[]User{
		{"tom","Password123"},
		{"jerry","Abc@123#"},
	}}

	fmt.Println(p)

}
