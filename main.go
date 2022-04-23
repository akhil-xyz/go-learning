package main

import "log"

func main() {
	log.Println("Hi there!")
	var color string = "red"
	log.Println("Color before change is ", color)
	changeColor(&color)
	log.Println("Color after change is ", color)
}

func changeColor(colorPtr *string) {
	*colorPtr = "black"
}
