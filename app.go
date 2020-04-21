package main

import (
	"fmt"
	"math/rand"
	"ndc/go/golang-book/exercises/arrays"
	"ndc/go/golang-book/exercises/closures"
	"ndc/go/golang-book/exercises/concurrency"
	"ndc/go/golang-book/exercises/maps"
	"ndc/go/golang-book/exercises/panic"
	"ndc/go/golang-book/exercises/pointers"
	"ndc/go/golang-book/exercises/problems"
	"ndc/go/golang-book/exercises/recursion"
	"ndc/go/golang-book/exercises/slices"
	"ndc/go/golang-book/exercises/structs"
)

func main() {
	fmt.Println("----- math/rand -----")
	fmt.Println("Random generated int:", rand.Intn(100))
	fmt.Println()

	fmt.Println("----- arrays -----")
	arrays.Main()
	fmt.Println()

	fmt.Println("----- closures -----")
	closures.Main()
	fmt.Println()

	fmt.Println("----- concurrency -----")
	concurrency.Main()
	fmt.Println()

	fmt.Println("----- maps -----")
	maps.Main()
	fmt.Println()

	fmt.Println("----- panic -----")
	panic.Main()
	fmt.Println()

	fmt.Println("----- pointers -----")
	pointers.Main()
	fmt.Println()

	fmt.Println("----- problems -----")
	problems.Main()
	fmt.Println()

	fmt.Println("----- recursion -----")
	recursion.Main()
	fmt.Println()

	fmt.Println("----- slices -----")
	slices.Main()
	fmt.Println()

	fmt.Println("----- structs -----")
	structs.Main()
	fmt.Println()

	// src: https://golang.org/pkg/strings/#example_Map
	// fmt.Println(string('A' + ('T'-'A'+13)%26))
	// fmt.Println(string('a' + ('w'-'a'+13)%26))
	// fmt.Println(string('a' + ('a'-'a'+13)%26))
	// fmt.Println(string('a' + ('s'-'a'+13)%26))
}
