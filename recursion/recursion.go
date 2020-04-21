package recursion

// Factorial : calculates and returns the factorial of a parameter passed unsigned integer
func Factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * Factorial(x-1)
}
