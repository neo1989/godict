package main

import (
	"fmt"
	"godict/lib"
	//"os"
	//"os/signal"
)

func main() {

	var word string

	for {
		fmt.Print("Input the word > ")
		fmt.Scanln(&word)
		trans.Trans(word)
	}

}
