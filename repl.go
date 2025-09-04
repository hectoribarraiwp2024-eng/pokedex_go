package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hectoribarra2024-eng/pokedex_go/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
		}
		
	}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"pokedex": {
			name:		 "pokedex",
			description: "Displays all the pokemon in your pokedex",
			callback:  	 commandPokedex,
		},
		"inspect": {
			name:		 "inspect <pokemon>",
			description: "Displays information about a pokemon",
			callback:	 commandInspect,
		},
		"catch": {
			name: 		 "catch <pokemon>",
			description: "Catches a pokemon",
			callback: 	 commandCatch,
		},
		"explore": {
			name: 		"explore <location_name>",
			description:"Displays more information about a specific location",
			callback:	commandExplore,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world.",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations.",
			callback:    commandMapb,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

/*func getConfig config{} {

}*/


