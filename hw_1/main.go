package main

import "fmt"

func main() {
	const s string = "abc2d10e" // testing string
	var l string = string(s[0]) //storage for the current letter
	var n string = "_"          //storage for the current number

	for i := 0; i < len(s); i++ {

		/* if this is the last element, then s[i+1] will break the program,
		so another algorithm is used for this iteration*/

		if i == len(s)-1 {
			if IsLetter(string(s[i])) {
				fmt.Printf(string(s[i]))
			} else {
				if IsNumeral(string(s[i])) {
					n += string(s[i])
					PrintLetter(StrToInt(n), l)
				} else {
					fmt.Printf("%v", string(s[i]))
				}
			}
			continue

			/* for all iterations except the last */
		} else {
			if IsLetter(string(s[i])) {
				l = string(s[i])
			} else {
				if IsNumeral(string(s[i])) {
					n += string(s[i])
				} else {
					fmt.Printf("%v", string(s[i]))
					continue
				}

			}
		}

		if IsLetter(string(s[i+1])) {
			PrintLetter(StrToInt(n), l)
			n = "_"
		}

	}

}
