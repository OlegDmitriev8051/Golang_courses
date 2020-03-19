package main

import (
	"fmt"
	f "github.com/OTUS/unpack_str/functions"
)

func main() {
	const s string = "sdas" // testing string
	var l string   //storage for the current letter
	var n string = "_"      /*storage for the current number.
							 "_" means that the number of prints of the lettet is 1 */

	for i := 0; i < len(s); i++ {

		/* if this is the last element, then s[i+1] will break the program,
		so another algorithm is used for this iteration*/
		if i == len(s)-1 {

			if !f.IsNumeral(string(s[i])) {
				l = string(s[i])
				f.PrintLetter(f.StrToInt(n), l)
				n = "_"
			} else {
				n += string(s[i])
				f.PrintLetter(f.StrToInt(n), l)
			}

			break
		
			/* for all iterations except the last */
		} else {

			if !f.IsNumeral(string(s[i])) {
				l = string(s[i])
			} else {
				n += string(s[i])
			}
		}

		/* if there is a letter after the digit,
		   then print the previous letter */
		if !f.IsNumeral(string(s[i+1])) {
			f.PrintLetter(f.StrToInt(n), l)
			n = "_"
		}

	}
	fmt.Println()
}
