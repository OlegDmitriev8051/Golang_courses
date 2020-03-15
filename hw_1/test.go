package main

import (
	"fmt"
	"math"
)

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

func main() {
	var n string = "_1231932032902"
	var n1 string = "_"

	var a int
	var b int

	a = StrToInt(n)
	b = StrToInt(n1)

	fmt.Println(a)
	fmt.Println(b)

}
