package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	repsponse, error := http.Get("htrestp://google.com")

	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}
	//fmt.Println(repsponse)

	// empty byte slice which have 99999 spaces
	// bs := make([]byte, 99999)
	// repsponse.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWriter{}

	// same as above 3 lines
	// func Copy(dst Writer, src Reader) (written int64, err error)
	io.Copy(lw, repsponse.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println(len(bs))
	return len(bs), nil
}
