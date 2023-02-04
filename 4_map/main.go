package main

import "fmt"

type colorList map[string]string

func main() {
	//var colors map[string]string

	//colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	colors["white"] = "#ffffff"

	delete(colors, "red")

	fmt.Println(colors)
	printMap(colors)

	myColors := colorList{
		"black": "#000000",
	}
	myColors.print()
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex for color", color, "is", hex)
	}
}

func (c colorList) print() {
	printMap(map[string]string(c))
}
