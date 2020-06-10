package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	res, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}
	io.Copy(lw, res.Body)
}

// implement the Writer interface (https://golang.org/pkg/io/#Writer) to use
// instead of os.Stdout in our io.Copy() usage
func (lw logWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	fmt.Println("Total bytes processed:", len(p))

	return len(p), nil
}
