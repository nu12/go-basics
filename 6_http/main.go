package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type myLogger struct{}

func main() {

	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Manually creating the byte slice
	// r := make([]byte, 99999)
	// resp.Body.Read(r)
	// fmt.Println(string(r))

	// Copying from a Reader to a Writer
	//io.Copy(os.Stdout, resp.Body)

	// Copying from a Reader to a custom Writer
	ml := myLogger{}
	io.Copy(ml, resp.Body)
}

func (myLogger) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	fmt.Println("Body length:", len(p))
	return len(p), nil
}
