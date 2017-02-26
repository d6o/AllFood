package main

import "fmt"

// App holds the main resources from the execution.
type App struct {
	radius       int
	lat          string
	lng          string
	providerList []Food
}

// Initialize receives all configs and initialize the program.
func (a *App) Initialize(radius int, lat, lng string) {
	a.radius = radius
	a.lat = lat
	a.lng = lng
}

func (a *App) AddProvider(p Food) {
	a.providerList = append(a.providerList, p)
}

// Run binds the port and keeps listening for connections.
func (a *App) Run() {
	result := map[string]Place{}
	for _, v := range a.providerList {
		places, err := v.Search(a.radius, a.lat, a.lng)
		if err != nil {
			panic(err)
		}
		for _, p := range places {
			result[p.Name] = p
		}
	}

	fmt.Printf("|%50s|%12s|%12s|\n", "Name", "Latitude", "Longitude")
	for _, v := range result {
		fmt.Printf("|%50s|%-12.8f|%-12.8f|\n", v.Name, v.Lat, v.Lng)
	}
}
