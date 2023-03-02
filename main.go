package main

import (
	"fmt"
	g "go-code-generation/generated"
)

type Service interface {
	GetEncodedString(string) string
}

func main() {
	var s Service = &g.EncodingService{}
	fmt.Println(s.GetEncodedString("my_password"))
}
