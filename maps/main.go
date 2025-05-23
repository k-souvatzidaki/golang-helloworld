package main

import (
	"fmt"
)

func main() {

	colors := map[string]string {
		"red": "#ff0000", 
		"blue": "#jas8dg",
	}

	fmt.Println(colors)


	// insert and delete
	colors2 := make(map[string]string)
	colors2["white"] = "#ffffff"
	fmt.Println(colors2)

	delete(colors2, "white") 
	fmt.Println(colors2)

	// loop in function 
	printMap(colors)
}

func printMap (c map[string]string) {
	fmt.Println("Inside printMap")
	for color,hex := range c {
		fmt.Println(color,hex)
	}
}