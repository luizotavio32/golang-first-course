package main

import "fmt"

func main(){
	name := "Luiz"
	version := 1.1
	fmt.Println("Hi", name)
	fmt.Println("Program version:", version)

	fmt.Println("1 - Initialize Monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exit")

	var command int
	fmt.Scan(&command)


}
