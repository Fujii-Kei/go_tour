package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	const Full = false
	fmt.Println("Go rules?", Truth)
	fmt.Println("R u hungry?", Full)
}
