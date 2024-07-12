package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/dixonwille/wmenu/v5"
)

type Message struct {
	Text string `json:"text"`
}

type Crop struct {
	Name                             string `json:"name"`
	NumWeeksToSproutIndoors          string `json:"numWeeksToSproutIndoors"`
	WeeksRelToFrostDateStartOutdoors string `json:"weeksRelToFrostDateStartOutdoors"`
	TotalGrowthMonths                string `json:"totalGrowthMonths"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	message := Message{Text: "Hello, World! I like cake"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
	message := Message{Text: "Goodbye, time to brush your teeth"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func main() {
	fmt.Println("Starting main.go")
	// added proxy to package.json so any API requests to /api/* will be proxied to the go backend running on port 8080.
	http.HandleFunc("/api/hello", helloHandler)
	http.HandleFunc("/api/goodbye", goodbyeHandler)
	http.ListenAndServe(":8080", nil)

	// Read the entire JSON file
	byteValue, err := os.ReadFile("data.json")
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	// Unmarshal the JSON data into a slice of structs
	var crops []Crop
	if err := json.Unmarshal(byteValue, &crops); err != nil {
		log.Fatalf("Error unmarshalling JSON data: %v", err)
	}

	for _, crop := range crops {
		fmt.Printf("Name: %s \n Weeks Indoors: %s \n Rel To Frost: %s \n Growth Months: %s\n",
			crop.Name, crop.NumWeeksToSproutIndoors, crop.WeeksRelToFrostDateStartOutdoors, crop.TotalGrowthMonths)
	}

}
