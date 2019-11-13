package main

import "github.com/priscillashort/go-spaceadventure-walkthrough/internal/spaceadventure"

func main() {
	spaceadventure.Start(spaceadventure.PlanetarySystem{
		Name: "Solar System", 
		Planets: []spaceadventure.Planet{
			spaceadventure.Planet{
				Name: "Jupiter", Description: "Desert planet",
			},
		},
	})
}
