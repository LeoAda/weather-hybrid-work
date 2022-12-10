package main

import "fmt"

func ChooseLocation() Location {
	//Display locations and ask user to choose one, return the chosen location
	var s string
	fmt.Print("Type city name: ")
	fmt.Scan(&s)
	locations := SearchLocationsList(s)
	for index, element := range locations {
		fmt.Println(index, ". ", element)
	}
	var i int
	fmt.Print("Choose a location: ")
	fmt.Scan(&i)
	return locations[i]
}
