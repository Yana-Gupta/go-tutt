package main

import (
	"fmt"
	"io"
	"os"
)

type writer struct {
}

func main() {
	p := []string{}
	p = os.Args

	file, err := os.Open(p[1])

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	// fmt.Println(*file)

	// io.Copy(os.Stdout, file)

	// fmt.Println(n)

	w := writer{}

	io.Copy(w, file)

}

func (writer) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	return len(bs), nil
}
