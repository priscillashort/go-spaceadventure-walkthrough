package spaceadventure

import (
	"fmt"
	"math/rand"
	"time"
)

func Start(planetarySystem PlanetarySystem) {
	printWelcome(planetarySystem)
	printGreeting(responseToPrompt("What is your name?"))
	fmt.Println("Let's go on an adventure!")
	travel(promptForRandomOrSpecificDestination("Shall I randomly choose a planet for you to visit? (Y or N)"), planetarySystem)
}

func printWelcome(planetarySystem PlanetarySystem) {
	fmt.Printf("Welcome to the %s!\n", planetarySystem.Name)
	fmt.Printf("There are %d planets to explore.\n", planetarySystem.NumberOfPlanets())
}

func responseToPrompt(prompt string) (response string) {
	fmt.Println(prompt)
	fmt.Scan(&response)
	return 
}

func printGreeting(name string) {
	fmt.Printf("Nice to meet you, %s. My name is Eliza, I'm an old friend of Siri.\n", name)
}


func travel(randomDestination bool, planetarySystem PlanetarySystem) {
	if (randomDestination) {
		travelToPlanet(selectRandPlanet(planetarySystem).Name)
	} else {
		travelToPlanet(responseToPrompt("Name the planet you would like to visit."))
	}
}

func promptForRandomOrSpecificDestination(prompt string) bool {
	var choice string
	for choice != "Y" && choice != "N" {
		choice = responseToPrompt(prompt)
		if choice == "Y" {
			return true
		} else if choice == "N" {
			return false
		} else {
			fmt.Println("Sorry, I didn't get that.")
		}
	}
	return false
}

func travelToPlanet(planetName string) {
	fmt.Printf("Traveling to %s...\n", planetName)
	
	//This should actually lookup the planet name in the options
	//and print the prompt that goes with the given planet
	//But this solution is not complete 
	fmt.Printf("Arrived at %s. The large red spot appears ominous.\n", planetName)
}

func selectRandPlanet(ps PlanetarySystem) (planet Planet){
	//select random planet
	fmt.Println("Selecting a random planet...")
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	planet = ps.Planets[rand.Intn(ps.NumberOfPlanets())]
	return
}
