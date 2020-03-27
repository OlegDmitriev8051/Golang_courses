package main

import (
	"fmt"
	"strings"
	"sort"
)
// this function returns slice with lower symbols
func lowSlice (wSlice []string) []string {
	var lowSlice []string 
	for _, v := range wSlice {
		lowSlice = append(lowSlice, strings.ToLower(v))
	} 
	return lowSlice
}

func main() {

	//key - word, value - amount of this word
	words := map[string]int{}

	//text constant is translating to slice of words
	var wSlice []string = strings.Split(text, " ")
	fmt.Println(wSlice)

	// make the low elements slice
	lowSlice := lowSlice(wSlice)
	fmt.Println(lowSlice)

	for _, v := range lowSlice {
		if _, ok := words[v]; !ok {
			words[v] = 1
		} else{
			words[v]++
		}
	}
	fmt.Println(words)

	// To sort in descending order,the map is written to the structure 
	type keyValue struct {
		key string
		value int	
	}
	var sortedStruct []keyValue
	for key, value := range words {
		sortedStruct = append(sortedStruct, keyValue{key, value})
	}
	// sorting structure
	sort.Slice(sortedStruct,func(i, j int) bool {
		return sortedStruct[i].value > sortedStruct[j].value
	})

	for _, k := range sortedStruct {
		fmt.Printf("%s, %d\n",k.key, k.value)
	}

}