// adapted from: https://withcodeexample.com/golang-tutorial-for-beginners/
package main

import "fmt"

func main() {
	// Declaration and creation of a map
	// Declaration of a map with string keys and string values
	var colors map[string]string
	// The make function is used to create an empty map
	colors = make(map[string]string)
	// Adding key-value pairs to the map
	colors["red"] = "#ff0000"
	colors["green"] = "#00ff00"
	colors["blue"] = "#0000ff"
	fmt.Println("Colors:", colors)

	// Adding key-value pairs
	moreColors := make(map[string]string)
	moreColors["red"] = "#ff0000"
	moreColors["green"] = "#00ff00"
	moreColors["blue"] = "#0000ff"

	// Accessing and modifying map elements
	redHex := moreColors["red"]
	moreColors["green"] = "#008000"
	fmt.Println("Red Hex:", redHex)
	fmt.Println("Updated Colors:", moreColors)

	// Deleting map elements
	otherColors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	delete(otherColors, "green")
	fmt.Println("Updated Colors:", otherColors)

	// Checking for the existence of a key
	colorsToCheck := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	if value, exists := colorsToCheck["green"]; exists {
		fmt.Println("Green Hex:", value)
	} else {
		fmt.Println("Green key does not exist.")
	}

	// Iterating over map elements
	for key, value := range colors {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// Map with non-string keys
	ages := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 35,
	}
	fmt.Println("Ages:", ages)

}
