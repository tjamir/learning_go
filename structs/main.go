package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newName string) {
	p.firstName = newName
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson",
		contactInfo: contactInfo{email: "alex@gmail.com", zip: 538438}}
	var jamir person
	jamir.firstName = "Thiago"
	jamir.lastName = "Jamir"
	jamir.contactInfo.email = "jamir@gmail.com"
	jamir.contactInfo.zip = 51011020
	alex.updateName("Lex")
	alex.print()
	jamir.print()
}
