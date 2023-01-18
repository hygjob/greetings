package main

import (
	"fmt"
	"log"
	//"github.com/hygjob/greetings/single"
	"github.com/hygjob/greetings/v2/single"
)

func main() {
	fmt.Println("hello ==> main()")
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := single.Hello("", true)
	if err == nil {
		fmt.Println(message)	
	} else {
		log.Fatal(err)
	}

}

