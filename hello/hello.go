package main

import (
	"fmt"
	"log"
	"github.com/hygjob/greetings/single"
)

func main() {
	fmt.Println("hello ==> main()")
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := single.Hello("John", true)
	if err == nil {
		fmt.Println(message)
	}

}

