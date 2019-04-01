package main

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "windows"
}

func main() {
	switch os := getOsName(); os {
	case "mac":
		fmt.Println("MAC")
	case "windows":
		fmt.Println("windows")
		// default:
		// 	fmt.Println("default")
	}

	t := time.Now()

	fmt.Println(t.Hour())

	switch {
	case t.Hour() > 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("Afternoon")
	}
}
