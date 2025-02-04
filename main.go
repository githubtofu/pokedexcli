package main
import (
	"fmt"
    "strings"
    "bufio"
    "os"
    "github.com/mtslzr/pokeapi-go"
//    "reflect"
)

type cliCommand struct {
    name    string
    description string
    callback func() error
}

var climap = map[string]cliCommand{}

func cleanInput(text string) []string {
    lower := strings.ToLower(text)
    words := strings.Fields(lower)
    return words
}	

func commandMap() error {
    l, err := pokeapi.Resource("location-area")
//    fmt.Printf("First Type:%v, Second Type:%v\n", reflect.TypeOf(l), reflect.TypeOf(l1))
    fmt.Println(l)
    return err
}

func commandExit() error {
    fmt.Printf("Closing the Pokedex... Goodbye!\n")
    os.Exit(0)
    return nil
}

func commandHelp() error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    fmt.Println("")
    for k, v := range climap {
        fmt.Printf("%v: %v\n", k, v.description)
    }
    return nil
}

func set_map() {
    climap["exit"] = cliCommand{
            name:   "exit",
            description: "Exit the Pokedex",
            callback: commandExit,
        }
    climap["help"] = cliCommand{
            name:   "help",
            description: "Displays a help message",
            callback: commandHelp,
        }
    climap["map"] = cliCommand{
            name:   "map",
            description: "Displays the names of 20 location areas in the Pokemon world",
            callback: commandMap,
        }
}

func main() {
    set_map()
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("Pokedex >")
        if scanner.Scan() {
            text := scanner.Text()
            tokens := cleanInput(text)
            if len(tokens) == 0 {
                fmt.Printf("Enter something\n")
                continue
            }
            for k, v := range climap {
                if k == tokens[0] {
                    err := v.callback()
                    if err != nil {
                        fmt.Println(err)
                    }
                }
            }
        } else {
            fmt.Printf("Error\n")
            break
        }
    }
}
