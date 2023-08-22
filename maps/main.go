package main

import "fmt"

func iterateMap(m map[string]string) {
	for k, v := range m {
		fmt.Println("Hex code for", k, "is", v)
	}
}

func main() {

	megacolors := make(map[string]string)

	megacolors["white"] = "ffffff"

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	iterateMap(colors)

	fmt.Println(colors)

}
