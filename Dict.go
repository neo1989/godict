package main

import (
	"fmt"
	"godict/lib"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nLeaving...")
		os.Exit(0)
	}()

	fmt.Println("\n======= Translator =======\n")
	for {
		var word string
		fmt.Print("Input the word > ")
		fmt.Scanln(&word)
		trans.Trans(word)
	}

}
