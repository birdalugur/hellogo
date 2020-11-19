package main

import (
	"fmt"

	"github.com/birdalugureren/greetings"

	"log"

	"example.com/user/hello/morestrings"
	"example.com/user/hello/randomgreeting"
)

func main() {
	message := greetings.Hello("uğur")
	fmt.Println(message)

	reversemessage, err := morestrings.ReverseRunes("Merhaba Dünya")
	fmt.Println(reversemessage)

	if err != nil {
		log.Fatal(err)
	}

	randomMessage, err := randomgreeting.Hello("Uğur")

	if err == nil {
		fmt.Println(randomMessage)
	}

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := randomgreeting.Hellos(names)

	if err != nil {
        log.Fatal(err)
    }
    // If no error was returned, print the returned map of
    // messages to the console.
    fmt.Println(messages)

}
