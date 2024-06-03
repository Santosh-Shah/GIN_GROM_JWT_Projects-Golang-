package main

import "fmt"

func main() {
	//var myMap map[string]int
	//
	//myMap = make(map[string]int)

	// adding key-value pairs to the map
	//myMap["Santosh"] = 25
	//myMap["Hariom"] = 30

	// Printing the value of map
	//fmt.Println("length:", len(myMap))
	//for _, value := range myMap {
	//	fmt.Println(value)
	//}

	//for key, _ := range myMap {
	//	fmt.Println(key)
	//}

	//value, exists := myMap["Santosh"]
	//fmt.Println(value)
	//fmt.Println(exists)
	//fmt.Println(myMap["Santosh"])

	//var mapDemo map[string]string
	//
	//mapDemo = make(map[string]string)
	//mapDemo["name1"] = "Santosh Shah"
	//mapDemo["name2"] = "hariom shah"
	//mapDemo["name3"] = "Rohit"
	//
	//fmt.Println(mapDemo["name1"])

	myMap := map[string]string{
		"name":    "Santosh Shah",
		"college": "Vedas College",
		"address": "Lalitpur, Nepal",
	}

	myMap["prevCollege"] = "National Infotech College"

	for _, value := range myMap {
		//fmt.Println(key, value)
		//fmt.Println(key)
		fmt.Println(value)

	}
}
