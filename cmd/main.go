package main

import (
	"fmt"
	"os"

	"github.com/brian-ding/swago/internal/pkg/swagger"
)

func main() {
	fmt.Println("hello, world!")
	// get the content from the file at root folder, with the path "./petstore.yaml"
	content, err := os.ReadFile("./petstore.yaml")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	swagger.Parse(content)
}
