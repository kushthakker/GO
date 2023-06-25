package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
	fmt.Println(resp, "resp")
	fmt.Printf("type is %T\n", resp)

	// bs := make([]byte, 999999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	io.Copy(os.Stdout, resp.Body)
}