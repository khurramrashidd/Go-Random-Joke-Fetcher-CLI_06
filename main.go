package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Joke struct defines the shape of the JSON data we expect to receive
type Joke struct {
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

func main() {
	fmt.Println("😂 Fetching a random joke for you from the internet...")

	// 1. Define the API endpoint
	url := "https://official-joke-api.appspot.com/random_joke"

	// 2. Make the HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching joke: %v\n", err)
		os.Exit(1)
	}
	// Defer closing the response body until the function finishes
	defer response.Body.Close()

	// 3. Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response data: %v\n", err)
		os.Exit(1)
	}

	// 4. Parse (Unmarshal) the JSON byte data into our Joke struct
	var joke Joke
	err = json.Unmarshal(body, &joke)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	// 5. Print the formatted joke
	fmt.Println("--------------------------------------------------")
	fmt.Printf("🤔 Setup:     %s\n", joke.Setup)
	fmt.Printf("💡 Punchline: %s\n", joke.Punchline)
	fmt.Println("--------------------------------------------------")
}