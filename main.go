package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	repsponse, error := http.Get("http://google.com")

	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}
	fmt.Println(repsponse)
}
