package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	helloMessage := stringutil.Reverse("Hello, OTUS!")
	fmt.Println(helloMessage)
}
