package spaceadventure

import "fmt"

func Start(planetarySystem PlanetarySystem) {
	printWelcome(planetarySystem)
	printGreeting(responseToPrompt("What is your name?"))
	fmt.Println("Let's go on an adventure!")
	travel(promptForRandomOrSpecificDestination("Shall I randomly choose a planet for you to visit? (Y or N)"))
}

func printWelcome(planetarySystem PlanetarySystem) {
	fmt.Printf("Welcome to the %s!", planetarySystem.Name)
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


func travel(randomDestination bool) {
	if (randomDestination) {
		travelToPlanet("")
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
	if planetName == "" {
		//This should actually generate a random planet name 
		//but the solution is not complete
		planetName = "Jupiter"
	}

	fmt.Printf("Traveling to %s...\n", planetName)
	
	//This should actually lookup the planet name in the options
	//and print the prompt that goes with the given planet
	//But this solution is not complete 
	fmt.Printf("Arrived at %s. The large red spot appears ominous.\n", planetName)
}
