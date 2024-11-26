package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = []rune("resume")
	var indexed = myString[0]

	fmt.Printf("%v, %T", indexed, indexed) //%T prints the type of something

	//string building

	var strSlice = []string{"t", "e", "s", "t"}
	// var builtStr = ""
	// for i := range strSlice {
	// 	builtStr += strSlice[i]
	// }
	// fmt.Printf("\n%v", builtStr)
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var builtString = strBuilder.String()
	fmt.Printf("\n%v", builtString)
	
}
