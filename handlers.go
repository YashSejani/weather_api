package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func handleWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "City parameter required", http.StatusBadRequest)
		return
	}

	// Check Redis Cache
	cachedData, err := rdb.Get(ctx, city).Result()
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(cachedData))
		return
	}

	// Fetch from External API on Cache Miss
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

	// Decode Response Stream
	var externalData ExternalWeatherResponse
	
	err = json.NewDecoder(resp.Body).Decode(&externalData)
	if err != nil {
		http.Error(w, "Failed to parse weather data", http.StatusInternalServerError)
		return
	}

	// Map to Clean Data Format
	cleanWeatherData := WeatherData{
		City:        externalData.ResolvedAddress,
		Temperature: fmt.Sprintf("%.1f°C", externalData.CurrentConditions.Temp),
		Condition:   externalData.CurrentConditions.Conditions,
		Humidity:    fmt.Sprintf("%.0f%%", externalData.CurrentConditions.Humidity),
		WindSpeed:   fmt.Sprintf("%.1f km/h", externalData.CurrentConditions.Windspeed),
		Sunrise:     externalData.CurrentConditions.Sunrise,
		Sunset:      externalData.CurrentConditions.Sunset,
		Icon:        externalData.CurrentConditions.Icon,
	}

	// Store in Cache
	cacheBytes, err := json.Marshal(cleanWeatherData)
	if err == nil {
		rdb.Set(ctx, city, cacheBytes, 10*time.Minute)
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cleanWeatherData)
}