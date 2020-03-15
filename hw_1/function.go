package main

import (
	//"math"
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

// func StrToInt(str string) int {
// 	var result int
// 	for i := 0; i < len(str); i++ {
// 		result += CharToInt(str[i]) * math.Pow10(len(str)-i-1)
// 	}
// }

// func CharToInt (ch string) int {
// 	if ch== {
// 		return 1
// 	} else {
// 		return int(ch - '0')
// 	}
// }
