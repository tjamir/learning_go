package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "4bf745",
	}

	var colors2 map[string]string

	colors3 := make(map[string]string)

	colors3["white"] = "#ffffff"

	fmt.Println(colors)
	fmt.Println(colors2)
	fmt.Println(colors3)

	delete(colors3, "white")
	delete(colors, "blue")
	fmt.Println(colors)
	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Color:", color, "Hex:", hex)
	}
}
