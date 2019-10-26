package main

import (
	"fmt"
	"reflect"

	"github.com/ngaut/log"
	tyaml "github.com/tidbops/tim/pkg/yaml"
	"gopkg.in/mikefarah/yaml.v2"
)

func main() {
	log.SetLevelByString("debug")

	yaml.DefaultMapType = reflect.TypeOf(yaml.MapSlice{})
	output, err := tyaml.Delete("./test.yml", "g")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}
