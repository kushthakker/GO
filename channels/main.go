package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string {
		"http://google.com",
		"http://amazon.in",
		"http://apple.com",
	}

	c := make(chan string)

	for _, url := range links {
		go check(url,c)
	
	}

	// for {
	// 	go check(<-c, c)
	// } // infinite loop

	// for i:=0; i<len(links); i++ {
	// 	go check(<-c, c)
	// }

	for l := range c{
		go func(link string) {
			time.Sleep(time.Second*5)
			check(link, c)
			}(l)
	} //  channel loop


}

func check(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Println(url, "is not up")
		c <- url
		return
	}
	c <- url
	fmt.Println(url, "is up")

}