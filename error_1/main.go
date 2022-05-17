package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//panic interception
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recovered", v)
		}
	}()

	file, err := os.Create("testFile.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer func() {
		closeErr := file.Close()
		if err == nil {
			err = closeErr
		}
	}()

	text := "test text"

	n, err := file.WriteString(text)
	if err != nil {
		panic(err)
	}

	fContent, err := ioutil.ReadFile("testFile.txt")
	if err != nil {
		panic(err)
	}
	//new error
	err = errors.New("test New error")
	//trigger a new error
	if string(fContent) != "1test text" {
		fmt.Println(err)
	}
	//referencing a non-existent array element
	c := fContent[99]

	fmt.Println(string(c))
	fmt.Printf("n: %d\n", n)
}
