package main

import "fmt"

func main() {
	const str string = "abc2d10e" // testing string
	var l string                  //storage for the current letter
	var n string                  //storage for the current number

	if !IsLetter(string(str[0])) {
		fmt.Println("Error!First symbol isn`t correct")
	}

	l = string(str[0])

	for i := 1; i < len(str); i++ {

		if IsLetter(string(str[i])) {
			// fmt.Println(l, n, "time(s)")
			// n = ""
			l = string(str[i])
		} else {
			if IsNumeral(string(str[i])) {
				n += string(str[i])
			}
		}

		// if IsLetter(string(str[i+1])) {
		// 	fmt.Println(n, "*", l)
		// 	n = ""
		// }
	}

}
