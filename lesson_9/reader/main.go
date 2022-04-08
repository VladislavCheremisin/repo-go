package main

import (
	"fmt"
	"lesson_9/reader/encoding/json"
	"lesson_9/reader/encoding/yaml"
)

func main() {
	fmt.Printf("%+v\n", json.ConfigReaderJson())
	fmt.Printf("%+v\n", yaml.ConfigReaderYaml())
}
