package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode string
}

type person struct {
	firstname string
	lastname  string
}

type personWithContact struct {
	firstname string
	lastname  string
	contact   contactInfo
}

func main() {
	// Way 1
	// Assingning in the ascending order
	p := person{"Yana", "Gupta"}
	fmt.Println(p.firstname, p.lastname)

	// Way 2
	p1 := person{firstname: "Alex", lastname: "Chunia"}
	fmt.Println(p1)

	// Way 3
	var p2 person
	// A different print statement
	fmt.Printf("%+v\n", p2)
	p2.firstname = "Mike"
	p2.lastname = "Tyson"
	fmt.Println(p2)

	alex := personWithContact{
		firstname: "Alex",
		lastname:  "Chunia",
		contact: contactInfo{
			email:   "alex@gmail.com",
			zipcode: "341998",
		},
	}

	alex.print()

	p.updateName("Beyoncé")
	fmt.Printf("%+v\n", p)

	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	update(mySlice)
	fmt.Println(mySlice[0])

	

// Go has pointer like C++ / C
// Value + with & = address
// address + with = value
func (p *person) updateName(name string) {
	p.firstname = name
	// (*p).firstname = name
}

func (p personWithContact) print() {
	fmt.Printf("%+v\n", p)
}

func update(s []string) {
	s[0] = "This is nothing"
}
