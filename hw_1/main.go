package main

func main() {
	
str := "a4bc2d10e"	
var slc [] string
var ch string
var q int // quantity of repeated chars

// if incpection first symbol	
for i,v := range str {
	if (str [i] < '0') || (str [i] > '9') { // if str[i] is not a number
		if i != 1 {// if it isn`t first iterration
			slc := append(slc, ch)
		}
		ch = v
	}else {
		q += v
	}
}	
} 
	
	
	
	
	
	
	
	/*type pair struct{
		i string
		s string
	}
	test := []pair{
		{"a4bc2d5e","aaaabccddddde"},
		{"abcd","abcd"},
		{"45",""},

	}
	for _, t := range test {
		if t.s == !!!(t.i) {
		   fmt.Printf("%d - %s\n", t.i, "OK")
		} else {
		   fmt.Printf("%d - %s\n", t.i, "FAIL")
		}
	 }
 }*/
}