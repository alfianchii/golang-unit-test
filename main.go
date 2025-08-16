package main

import (
	"fmt"

	"github.com/alfianchii/golang-unit-test/helper"
)

func main() {
	var result string = helper.HelloWorld("OK")
	fmt.Println(result)
}