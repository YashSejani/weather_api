package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	// "io"
	"github.com/joho/godotenv"
)

type WeatherData struct {
	City        string `json:"city"`
	Temperature string `json:"temperature"`
	Condition   string `json:"condition"`
}

type ExternalWeatherResponse struct {
	ResolvedAddress   string `json:"resolvedAddress"`
	CurrentConditions struct {
		Temp       float64 `json:"temp"`
		Conditions string  `json:"conditions"`
	} `json:"currentConditions"`
}

func weather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")

	if city == "" {
		http.Error(w, "City perameter required", http.StatusBadRequest)
		return
	}

	apiURL := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s/today?unitGroup=metric&key=%s", city, os.Getenv("WEATHER_API_KEY"))

	resp, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, "Failed to fetch weather data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "External weather service error or invalid city", resp.StatusCode)
		return
	}

	var externalData ExternalWeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&externalData)
	if err != nil {
		http.Error(w, "Failed to parse weather data data", http.StatusInternalServerError)
		return
	}

	cleanWeatherData := WeatherData{
		City:        externalData.ResolvedAddress,
		Temperature: fmt.Sprintf("%.1f°C", externalData.CurrentConditions.Temp),
		Condition:   externalData.CurrentConditions.Conditions,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cleanWeatherData)

	// bodyBytes, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	http.Error(w, "Failed to read response body", http.StatusInternalServerError)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")

	// w.Write(bodyBytes)

	// fmt.Println("--- FULL API RESPONSE START ---")
	// fmt.Println(string(bodyBytes))
	// fmt.Println("--- FULL API RESPONSE END ---")
}

func main() {
	err := godotenv.Load()
    if err != nil {
        fmt.Println("Warning: No .env file found, relying on system env")
    }

	mux := http.NewServeMux()

	mux.HandleFunc("/weather", weather)

	http.ListenAndServe(":8080", mux)
}
