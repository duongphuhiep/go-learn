package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Weather represents weather data for a city
type Weather struct {
	City      string  `json:"city"`
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	Condition string  `json:"condition"`
}

// inMemoryWeather stores fake weather data for some cities
var inMemoryWeather = map[string]Weather{
	"London":   {City: "London", Temp: 15.2, FeelsLike: 12.5, Condition: "Cloudy"},
	"New York": {City: "New York", Temp: 10.8, FeelsLike: 8.3, Condition: "Sunny"},
	"Tokyo":    {City: "Tokyo", Temp: 22.1, FeelsLike: 20.4, Condition: "Rainy"},
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/weather/{city}", func(w http.ResponseWriter, r *http.Request) {
		city := r.PathValue("city")

		weather, ok := inMemoryWeather[city]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "City '%s' not found", city)
			return
		}

		// Encode weather data to JSON
		data, err := json.Marshal(weather)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error marshalling weather data: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Server listening on port 8080")
	server.ListenAndServe()
}
