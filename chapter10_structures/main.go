package main
import "fmt"


// we defined a  struct of the type customer
type Customer struct {
	name string
	username string
	password string
	age uint
}


// this is termed as  a nested  struct
/*
here we  are using one custom type
and using  it  as an attribute for another
*/
type Product struct {
	productId int
	customers []Customer
}


//acccessing  individual properties  ->  using the . operator
func printCustomerDetails(customer Customer){
	fmt.Printf(" Name:%s \n username:%s \n password:%s  \n  age:%d\n",customer.name,customer.username,customer.password,customer.age)
	// notice that we  are using the . operator to access the individual properties of customer
}

func main(){
// declaring a struct
var c Customer  //  all the struct  fields are initialized to 0
fmt.Println(c)

//declaring and initializing
var c1 Customer =  Customer{"itachi","realsusanoo","youareundermygenjutsu",21}
fmt.Println(c1)

//short hand notation
c2:=Customer{"Naruto","ramen","dattebayo",17}
printCustomerDetails(c2)



//  using the  nested struct
product  := Product {
1, 
[]Customer{
{"Santadip","rudra","abc",21},
{"Ash","kethum","bcd",34},},
}

fmt.Println(product)

}
