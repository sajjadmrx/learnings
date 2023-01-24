package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var pl = fmt.Println

func main() {

	pl("whats your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	} else {
		pl("Hello Dear", name, err)
	}

}
