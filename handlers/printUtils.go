package handlers

import (
	"fmt"
)

func PrintHeader(str string) {
	fmt.Println("***********************************************************************")
	l := len(str)
	l = 67 - l
	if l%2 != 0 {
		str = str + " "
	}
	l = l / 2
	for i := 0; i < l; i++ {
		str = " " + str + " "
	}
	str = "**" + str + "**"
	fmt.Println(str)
	fmt.Println("***********************************************************************")
}
