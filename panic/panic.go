package panic

import "fmt"

// DeferredRecovery : declares a deferred, nested recovery function and then calls the panic function
func DeferredRecovery() {
	/*
		The "defer" keyword tells Go to execute the deferred function (call) after the rest of the function
		completes (i.e. call "panic" function then call the nested function), can be used in non-panic scenarios
		The "panic" function causes a runtime error and stops execution of the goroutine
		The "recover" function stops the panic and returns the value that was passed to the panic (i.e. "Oops I did it again...")
	*/
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("Oops I did it again...")
}
