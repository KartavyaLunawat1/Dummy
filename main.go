package main

import "fmt"

func printNumb(Data int) {
	if Data == 0 {
		return
	}
	fmt.Println(Data)
	printNumb(Data - 1)
}

func main() {
	printNumb(5)
}
