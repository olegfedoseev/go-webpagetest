package main

import (
	"fmt"
	"github.com/olegfedoseev/webpagetest"
	"os"
)

func main() {
	fmt.Println("WebPageTest")

	if len(os.Args) == 0 {
		fmt.Println("You should provide URL for test!")
		os.Exit(2)
	}

	fmt.Println("arg", os.Args[0])

	wpt, _ := webpagetest.NewClient("http://webpagetest.app.s")

	err := wpt.CancelTest("161124_QZ_1")
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(2)
	}

	// Get Testers
	// result, err := wpt.GetTesters()
	// if err != nil {
	// 	fmt.Printf("Error: %v", err)
	// 	os.Exit(2)
	// }
	// fmt.Printf("Result: %#v", result)

	// Get Locations
	// result, err := wpt.GetLocations()
	// if err != nil {
	// 	fmt.Printf("Error: %v", err)
	// 	os.Exit(2)
	// }
	// fmt.Printf("Result: %#v", result)

	// Test Status
	// result, err := wpt.GetTestStatus("161122_K9_A")
	// if err != nil {
	// 	fmt.Printf("Error: %v", err)
	// 	os.Exit(2)
	// }
	// fmt.Printf("Result: %#v", result)

	// result, err := wpt.RunTestAndWait("http://ya.ru")
	// if err != nil {
	// 	fmt.Printf("Error: %v", err)
	// 	os.Exit(2)
	// }

	// fmt.Printf("Result: %#v", result)
}
