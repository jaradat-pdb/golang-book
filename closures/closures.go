package closures

// MakeEvenGenerator : returns a function which generates even numbers
func MakeEvenGenerator() func() uint {
	/*
		A closure comprises of an embedded function (i.e. a function within
		a function) and the non-local variable(s) referenced by the embedded
		function (i.e. the variable i, of unsigned integer type, below)
		Such variables persist between calls, thus the first call to this function
		will return ret = 0 and increment i by 2, the next function call
		will return ret = 2 and increment i by 2, and so on...
	*/
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// MakeOddGenerator : returns a function which generates odd numbers
func MakeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
