package functions

import (
	"fmt"
	"math"
	"strings"
)

func IsLetter(s string) bool {

	const symb string = "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtVvUuWwXxYyZz"

	if strings.ContainsAny(symb, s) {
		return true
	} else {
		return false
	}
}

func IsNumeral(s string) bool {

	const num string = "0123456789"

	if strings.ContainsAny(num, s) {
		return true
	} else {
		return false
	}
}

func StrToInt(str string) int {

	var result int

	if str == "_" {
		return 1
	} else {
		slice := []byte(str)
		slice = append(slice[:0], slice[1:]...) // delete first element
		for i := 0; i < len(slice); i++ {
			result += CharToInt(slice[i]) * int(math.Pow10(len(slice)-i-1))
		}
		return result
	}

}

func CharToInt(ch byte) int {
	return int(ch - '0')
}

func PrintLetter(n int, l string) {
	for i := 0; i < n; i++ {
		fmt.Printf("%v", l)
	}
}
