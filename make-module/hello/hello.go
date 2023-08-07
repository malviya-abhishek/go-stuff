package main
import (
	"fmt"
	"example.com/greetings"
	"log"
)

func main(){
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// message, err := greetings.Hello("Joe")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(message)

	names := []string {"Joe", "Larry", "Tom"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)


}