package main

import "fmt"

func main() {
	colors := make(map[string]string)

	// All the keys inside this braces are tight
	// Means all the keys are string here
	// You can't do colors.white

	colors["white"] = "#fff"
	colors["black"] = "#000000"

	// delete(colors, "white")
	// delete(colors, "black")

	fmt.Println(colors)

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
