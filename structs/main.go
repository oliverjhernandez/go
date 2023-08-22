package main

import "fmt"

type person struct {
	firstname   string
	lastname    string
	contactInfo contactInfo
}

type contactInfo struct {
	email   string
	zipcode string
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// Next line: *person is a type reference and it means a pointer to a person type
func (personPointer *person) updateName(newFirstName string) {
	// Next line: *personPointer refers to the data stored in personPointer
	(*personPointer).firstname = newFirstName
}

func main() {

	alex := person{
		firstname: "Alex",
		lastname:  "Anderson"}

	// Next line: &alex refers to alex memory address or pointer
	alexPointer := &alex
	alexPointer.updateName("Carl")
	// alex.print()

	var oliver person
	oliver.firstname = "Oliver"
	oliver.lastname = "Hernandez"
	oliver.print()
	oliverPointer := "what"
	fmt.Println(oliverPointer)

	chuck := person{
		firstname: "Chuck",
		lastname:  "Norris",
		contactInfo: contactInfo{
			email:   "chuck@power.com",
			zipcode: "999999",
		},
	}
	chuck.print()

}
