package spaceadventure

import "fmt"

func Start() {
	PrintWelcome()
	PrintGreeting(GetResponseToPrompt("What is your name?"))
	fmt.Println("Let's go on an adventure!")
	Travel()
}

func PrintWelcome() {
	fmt.Println("Welcome to the Solar System!")
	fmt.Println("There are 8 planets to explore.")
}

func GetResponseToPrompt(prompt string) (response string) {
	fmt.Println(prompt)
	fmt.Scan(&response)
	return 
}

func PrintGreeting(name string) {
	fmt.Printf("Nice to meet you, %s. My name is Eliza, I'm an old friend of Siri.\n", name)
}

func Travel() {
	var choice string
	for choice != "Y" && choice != "N" {
		choice = GetResponseToPrompt("Shall I randomly choose a planet for you to visit? (Y or N)")
		if choice == "Y" {
			TravelToPlanet("")
		} else if choice == "N" {
			planetName := GetResponseToPrompt("Name the planet you would like to visit.")
			TravelToPlanet(planetName)
		} else {
			fmt.Println("Sorry, I didn't get that.")
		}
	}
}

func TravelToPlanet(planetName string) {
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
