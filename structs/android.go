package structs

import "fmt"

// Android : defines an android struct type with Model representing the android's model and an embedded Person type
type Android struct {
	Person
	Model string
}

// Talk : prints a hello message with the android's name and model
func (a *Android) Talk() {
	a.Person.Talk()
	fmt.Println("I am a", a.Model, "model")
}
