package main

import (
	"fmt"
	f "github.com/OTUS/unpack_str/functions"
)

func main() {
	const s string = "gJd8SA0!3" // testing string
	var l string = string(s[0])  //storage for the current letter
	var n string = "_"           //storage for the current number

	// if IsNumeral(string(s[0])) {

	// }
	for i := 0; i < len(s); i++ {

		/* if this is the last element, then s[i+1] will break the program,
		so another algorithm is used for this iteration*/

		if i == len(s)-1 {
			if !f.IsNumeral(string(s[i])) {
				fmt.Printf(string(s[i]))
			} else {
				if f.IsNumeral(string(s[i])) {
					n += string(s[i])
					f.PrintLetter(f.StrToInt(n), l)
				} else {
					fmt.Printf("%v", string(s[i]))
				}
			}
			continue

			/* for all iterations except the last */
		} else {
			if !f.IsNumeral(string(s[i])) {
				l = string(s[i])
			} else {
				if f.IsNumeral(string(s[i])) {
					n += string(s[i])
				} else {
					fmt.Printf("%v", string(s[i]))

				}

			}
		}

		if !f.IsNumeral(string(s[i+1])) {
			f.PrintLetter(f.StrToInt(n), l)
			n = "_"
		}

	}
	fmt.Println()
}
