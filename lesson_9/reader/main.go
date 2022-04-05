package main

import (
	"fmt"
	configurationReaderJSON "lesson_9/reader/configurationJSON"
	configurationReaderYAML "lesson_9/reader/configurationYAML"
)

func main() {
	fmt.Printf("%+v\n", configurationReaderJSON.ConfigReaderJSON())
	fmt.Printf("%+v\n", configurationReaderYAML.ConfigReaderYAML())
}
