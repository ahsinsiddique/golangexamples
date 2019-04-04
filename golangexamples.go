package golangexamples

import "github.com/ehteshamz/greetings"

// Task1
func ConcatSlice(sliceToConcat []byte) string {
	var str string // variable
	for var1 := 0; var1 < len(sliceToConcat); var1++ {
		str += string(sliceToConcat[var1])
		if var1 != len(sliceToConcat)-1 {
			str += "-"
		}
	}
	return str
}

// Task2
func Encrypt(sliceToEncrypt []byte, ceaserCount int) {
	for var1 := 0; var1 < len(sliceToEncrypt); var1++ {
		sliceToEncrypt[var1] = (sliceToEncrypt[var1]) + byte(ceaserCount)
	}
}

//Task3
func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)
}
