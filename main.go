package main

import (
	"fmt"
	"net/http"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: No .env file found, relying on system env")
	}

	initRedis()

	mux := http.NewServeMux()
	
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/", fileServer)
	
	mux.HandleFunc("/weather", handleWeather)

	http.ListenAndServe(":8080", mux)
}