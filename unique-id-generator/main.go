package main

import (
	"fmt"
	"unique-id-generator/idgenerator"
)

func main() {
	generator := idgenerator.NewCounterGenerator()
	for i := 0; i < 10; i++ {
		fmt.Println(generator.GenerateID())
	}
}
