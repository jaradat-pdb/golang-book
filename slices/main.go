package slices

import "fmt"

func Main() {
	fmt.Println(AppendSlice())
	arrX, arrY := CopySlice()
	fmt.Println("x[]:", arrX, "|| y[]:", arrY)
}
