package structs

import "fmt"

// Person : defines a person struct type with Name representing the name
type Person struct {
	Name string
}

// Talk : prints a hello message with the person's name
func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}
