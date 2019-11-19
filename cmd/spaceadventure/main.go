package main

import "github.com/priscillashort/go-spaceadventure-walkthrough/internal/spaceadventure"

func main() {
	spaceadventure.Start(spaceadventure.PlanetarySystem{
		Name: "Solar System", Planets: mockPlanets(),
	})
}

func mockPlanets() []spaceadventure.Planet {
	return []spaceadventure.Planet{
		spaceadventure.Planet{Name: "Canton", Description: "Mud planet",},
		spaceadventure.Planet{Name: "Dagobah", Description: "Swamp Planet"},
		spaceadventure.Planet{Name: "Hoth", Description: "Icy Planet"},
		spaceadventure.Planet{Name: "Tatooine", Description: "Desert Planet"},
		spaceadventure.Planet{Name: "Jakuu", Description: "Party Planet"},
		spaceadventure.Planet{Name: "Bespin", Description: "Cloud Planet"},
		spaceadventure.Planet{Name: "Endor", Description: "Forest Moon"},
		spaceadventure.Planet{Name: "Vogsphere", Description: "Bad Poetry"},
		spaceadventure.Planet{Name: "Coruscant", Description: "City Planet"},
		spaceadventure.Planet{Name: "Gallifrey", Description: "Timelord Planet"},
	}
}
