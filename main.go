package main

import (
	"fmt"
	"log"

	"github.com/PJNube/universal-extension/lib"
)

func main() {
	err := lib.PackageExtension()
	if err != nil {
		log.Fatalf("Error packaging extension: %v", err)
	} else {
		fmt.Println("Extension packaged successfully!")
	}
}
