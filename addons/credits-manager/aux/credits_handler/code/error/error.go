package error

import (
	"fmt"
	"log"
)

func HandleErrorMessage(msg string, err error) {
	log.Fatal(fmt.Sprintf("%s: %s", msg, err.Error()))
}

func Message(msg string) {
	log.Println(msg)
}
