package main
import (
	"fmt"
    "strings"
    "bufio"
    "os"
    "github.com/mtslzr/pokeapi-go"
    "github.com/mtslzr/pokeapi-go/structs"
    "io"
    "encoding/json"
    "net/http"
)

type cliCommand struct {
    name    string
    description string
    callback func(*config) error
}

type config struct {
    Next string
    Previous string
}

var climap = map[string]cliCommand{}

func cleanInput(text string) []string {
    lower := strings.ToLower(text)
    words := strings.Fields(lower)
    return words
}	

func getResults(url string, c *config) ([]structs.Result, error) {
    results := []structs.Result{}
    res, err := http.Get(url)
    if err != nil {
        return results, err
    }
    defer res.Body.Close()
    body, err := io.ReadAll(res.Body)
    if res.StatusCode > 299 {
        return results, fmt.Errorf("Response failed with code : %d and \nbody: %s\n", res.StatusCode, body)
    }
    if err != nil {
        return results, err
    }
    resource := structs.Resource{}
    err = json.Unmarshal(body, &resource)
    if err != nil {
        return results, err
    }
    results = resource.Results
    c.Next = resource.Next
    if resource.Previous != nil {
        c.Previous = resource.Previous.(string)
    } else {
        c.Previous = ""
    }
    return results, nil
}

func commandMap(c *config) error {
    results := []structs.Result{}
    if c.Next == "" {
        l, err := pokeapi.Resource("location-area")
        if err != nil {
            return err
        }
        c.Next = l.Next
        if l.Previous != nil {
            c.Previous = string(l.Previous.(string))
        }
        results = l.Results
    } else {
        r, err := getResults(c.Next, c)
        if err != nil {
            return err
        }
        results = r
    }
    for _, a_result := range results {
        fmt.Println(a_result.Name)
    }
    return nil
}

func commandMapb(c *config) error {
    results := []structs.Result{}
    if c.Previous == "" {
        fmt.Println("you're on the first page")
        return nil
    } else {
        r, err := getResults(c.Previous, c)
        if err != nil {
            return err
        }
        results = r
    }
    for _, a_result := range results {
        fmt.Println(a_result.Name)
    }
    return nil
}

func commandExit(c *config) error {
    fmt.Printf("Closing the Pokedex... Goodbye!\n")
    os.Exit(0)
    return nil
}

func commandHelp(c *config) error {
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
    climap["mapb"] = cliCommand{
            name:   "mapb",
            description: "Displays the names of previous 20 location areas in the Pokemon world",
            callback: commandMapb,
        }
}

func main() {
    set_map()
    this_conf := config{}
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
                    err := v.callback(&this_conf)
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
