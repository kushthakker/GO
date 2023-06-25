package main

import "fmt" 

func main(){
	// var colors map[string]string
	// colors:= make(map[string]string)
	colors:= map[string]string{
		"red": "#FF0",
		"green": "#4bf",
	}
	colors["white"] = "#fff"
	fmt.Println(colors)
	printMain(colors)

}


func printMain(c map[string]string) {
	for key, value := range c {
		fmt.Println("key",key, "value",value)
	}
}