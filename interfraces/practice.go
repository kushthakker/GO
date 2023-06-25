package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]
	fmt.Println(fileName)
	f, err := os.OpenFile(fileName, os.O_RDWR, 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)

}