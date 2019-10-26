package main

import (
	"fmt"
	"log"

	"github.com/tidbops/tim/pkg/parser"
)

func main() {
	p := parser.NewParser()

	new, delete, err := p.ParserFile("../../rules/v2.1.17-to-v3.0.4.yml", ".", "tikv")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(new)
	fmt.Println(delete)
}
