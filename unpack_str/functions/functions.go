package functions

import (
	"fmt"
	"math"
	"strings"
)
// IsLetter() checks if a character is a letter. IsLetter is optional and not used in algorithm.
func IsLetter(s string) bool {

	const symb string = "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtVvUuWwXxYyZz"

	if strings.ContainsAny(symb, s) {
		return true
	} 
	return false
	}

// IsNumeral() checks if a character is a numeral.	
func IsNumeral(s string) bool {

	const num string = "0123456789"

	if strings.ContainsAny(num, s) {
		return true
	}
	return false
	
}
// StrToInt() converts a string to a number(int)
func StrToInt(str string) int {

	var result int

	if str == "_" {
		return 1
	} 

	slice := []byte(str)
	slice = append(slice[:0], slice[1:]...) // delete first element
	for i := 0; i < len(slice); i++ {
		result += CharToInt(slice[i]) * int(math.Pow10(len(slice)-i-1))
	}
	return result
}
// CharToInt() converts a char(byte) to a number(int)
func CharToInt(ch byte) int {
	return int(ch - '0')
}
// PrintLetter() prints a char a required number of times
func PrintLetter(n int, l string) {
	for i := 0; i < n; i++ {
		fmt.Printf("%v", l)
	}
}
