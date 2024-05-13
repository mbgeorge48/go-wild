package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

	names := []string{"gregory", "wallace", "mr whimma"}

	messages ,err := greetings.Hellos(names)
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

	for _, message := range messages{
		fmt.Println(message)
	}
}
