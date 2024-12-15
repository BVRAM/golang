package main

import "fmt"

func main() {
	fmt.Println("maps in golang")
	mapsGo()
}

func mapsGo() {

	// Basic syntax - Declaration -	var mapName map[KeyType]ValueType
	var mapName map[int]string
	fmt.Println(mapName)

	// Creating a Map:
	myMap := make(map[string]int)

	// Adding Elements:
	myMap["apple"] = 1
	myMap["banana"] = 2
	myMap["chocolate"] = 3
	myMap["dragon"] = 4

	// Deleting Elements
	delete(myMap, "banana")

	// Accessing Elements
	value, ok := myMap["egg"]

	if ok {
		fmt.Println("Value: ", value)
	} else {
		fmt.Println("Key not found")
	}

	// Iterating over a map
	for key, value := range myMap {
		fmt.Println(key, value)
	}

	// Map Capacity
	myMap1 := make(map[string]int, 10) // Pre-allocate a map with capacity 10
	myMap1["a"] = 1
	myMap1["b"] = 2
	myMap1["c"] = 3
	myMap1["d"] = 4

	// Iterating map
	for key, value := range myMap1 {
		fmt.Println(key, value)
	}

	// Deleting Multiple Elements
	keysToDelete := []string{"a", "b"}

	for _, key := range keysToDelete {
		delete(myMap1, key)
	}

	// Iterating map
	for key, value := range myMap1 {
		fmt.Println(key, value)
	}

	// Checking for Key Existence
	value, ok = myMap1["d"]
	if ok {
		fmt.Println("Key exists, use the value", value)
	} else {
		fmt.Println("Key doesn't exist")
	}

	// Creating Maps from Slices
	slice := []struct {
		Key   string
		Value int
	}{
		{"apple", 10},
		{"banana", 20},
	}

	myMap2 := make(map[string]int, len(slice))
	for _, pair := range slice {
		myMap2[pair.Key] = pair.Value
	}

}
