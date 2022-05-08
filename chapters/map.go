package main

import "fmt"

func main() {
	myMap := make(map[int]string) //make for initialising the map.
	myMap[1] = "akhil"
	myMap[2] = "xavier"
	myMap[3] = "paul"
	myMap[4] = "irene"
	myMap[5] = "ann"
	myMap[6] = "sumi"

	for _, item := range myMap {
		fmt.Println(item)
	}

	fmt.Println(myMap[4])
}
