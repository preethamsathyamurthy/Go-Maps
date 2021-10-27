package main

import "fmt"

func main() {
	// A map is similar to a struct but it can only have a predefined
	// type of key and value
	// in the below example - we are determining
	// the type of key to be string and value
	// to be string
	color := map[string]string{
		"Red":   "#ff0000",
		"Blue":  "#0000FF",
		"Green": "#00FF00",
	}

	//map values are accessed by [](square braces) mapping
	// struct values are accessed by .(dot) mapping
	color["purple"] = "800080"

	printMap(color)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("The hex value of ", color, " is ", hex)
	}
}
