package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func main() {
	resp, err := http.Get("https://example.com/")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// io.Copy takes two arguments
	l := logWriter{}
	io.Copy(l, resp.Body)
	fmt.Println("")
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	return len(bs), nil
}
