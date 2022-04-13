package main

import (
	"fmt"
	"time"
)

func main() {
	name := "Felix Anducho"
	otherNames := "Argel Puentes"
	isOtherNames := false
	today := time.Now()
	if isOtherNames {
		fmt.Println("Hi " + otherNames)
		fmt.Println("---------------------")
		fmt.Println("Today is", today.Format("01-02-2006"))
	} else {
		fmt.Println("Hi " + name)
		fmt.Println("---------------------")
		fmt.Println("Today is", today.Format("01-02-2006"))
	}
}
