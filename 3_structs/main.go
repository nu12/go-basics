package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo // Nested struct
}

func main() {
	john := person{"John", "Doe", contactInfo{"jd@email.com", "1234"}} // depends on the order of the fields
	jane := person{
		firstName: "Jane",
		lastName:  "Doe",
		contact: contactInfo{
			email:   "jjd@email.com",
			zipCode: "1234",
		},
	}

	var michael person
	michael.firstName = "Michael"
	michael.lastName = "Doe"

	john.updateName("Jimmy")
	john.print()
	jane.print()

}

func (pPtr *person) updateName(newFirstName string) {
	(*pPtr).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
