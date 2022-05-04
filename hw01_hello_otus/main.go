package main

import (
	"golang.org/x/example/stringutil"
	"io"
	"os"
)

func Reverse(s string) string {
	return stringutil.Reverse(s)
}

func main() {
	_, err := io.WriteString(os.Stdout, Reverse("Hello, OTUS!"))
	if err != nil {
		panic(err)
	}
}
