package main

import (
	"fmt"
	g "go-code-generation/generated"
)

type Service interface {
	GetGeneratedData(string) string
}

func main() {
	var s Service = &g.GenerationService{}
	fmt.Println(s.GetGeneratedData("my_password"))
}
