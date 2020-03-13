package main

import (
	"fmt"
	"strings"
)

func main() {

	const str string = "0aDV2a55DSD6xz7v809z"

	const symb string = "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtVvUuWwXxYyZz"

	const num string = "0123456789"

	var slc []string

	// transform strToInt
	// for _, v := range str {
	// 	fmt.Println(int(v - '0'))
	// }

	// isLetter
	// for _, v := range str {
	// 	if strings.ContainsAny(symb, string(v)) {
	// 		fmt.Printf("%v\n", string(v))
	// 	}
	// }

	// isNumerical
	for _, v := range str {
		if strings.ContainsAny(num, string(v)) {
			slc = append(slc, string(v))
		}
	}
	fmt.Printf("%v\n", slc)

}
