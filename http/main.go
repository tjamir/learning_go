package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	readHTMLHard()

}

func readHTMLEasy() {
	resp, error := http.Get("http://www.google.com")
	if error != nil {
		fmt.Println("Error", error)
		os.Exit(1)
	}
	io.Copy(os.Stdout, resp.Body)
}

func readHTMLHard() {
	resp, error := http.Get("http://www.google.com")
	if error != nil {
		fmt.Println("Error", error)
		os.Exit(1)
	}
	bytes := make([]byte, 0)
	//fmt.Println(resp)
	read := 1

	eof := false

	for read > 0 && !eof {
		temp := make([]byte, 128)
		n, error := resp.Body.Read(temp)
		if n < 128 {
			temp = temp[:n]
		}
		bytes = append(bytes, temp...)
		if error != nil {
			if error.Error() == "EOF" {
				eof = true
			} else {
				fmt.Println("Error", error)
				os.Exit(2)
			}
		}
		read = n

	}
	resp.Body.Close()
	fmt.Println(string(bytes))
}
