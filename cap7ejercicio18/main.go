package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"

)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	node, err := xmlnode.Parse(dec)
	if err != nil {
		log.Fatalf("ch07/ex18: %v", err)
	}
	fmt.Printf("%s\n", node)
}