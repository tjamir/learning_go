package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson",
		contact: contactInfo{email: "alex@gmail.com", zip: 538438}}
	var jamir person
	jamir.firstName = "Thiago"
	jamir.lastName = "Jamir"
	jamir.contact.email = "jamir@gmail.com"
	jamir.contact.zip = 51011020
	fmt.Println(alex)
	fmt.Printf("%+v", jamir)
}
