package main

import (
	"os"
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server successfully running on port :%s\n", port)
	http.ListenAndServe(":"+port, mux)
}