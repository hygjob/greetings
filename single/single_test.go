package single_test

import (
	"fmt"
	"log"
	"testing"
	"github.com/hygjob/greetings/single"
)

func TestMain(t * testing.T) {
	fmt.Println("hello ==> main()")
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := single.Hello("John", true)
	if err == nil {
		fmt.Println(message)
	}

}

