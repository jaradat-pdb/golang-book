package problems

// Halfer : takes an integer and halves and determines if the integer is even or odd
func Halfer(i int) (int, bool) {
	j := i / 2
	isEven := false
	if i%2 == 0 {
		isEven = true
	}
	return j, isEven
}

// GreatestInList : takes a variadic (variable) amount of integers and determines the greatest one in value
func GreatestInList(args ...int) int {
	greatest := 0
	for _, value := range args {
		if value > greatest {
			greatest = value
		}
	}
	return greatest
}

// FibonacciSequencer : calculates the fibonacci sequence of a parameter passed unsigned integer
func FibonacciSequencer(x uint) uint {
	switch x {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return (FibonacciSequencer(x-1) + FibonacciSequencer(x-2))
	}
}
