package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"flag"
)

func main() {
	var city string
	flag.StringVar(&city, "city", "", "City name for which to fetch weather data")
	flag.Parse()
  
	if city == "" {
	  fmt.Println("Please provide a city name using the -city flag")
	  return
	}

	weatherData, err := getWeatherCity(city)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Extract and print relevant information
	fmt.Println("---", city, "---")
	fmt.Printf("Current Temperature: %v°F\n", weatherData["main"].(map[string]interface{})["temp"])
}

func getWeatherCity(city string) (map[string]interface{}, error) {
	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	// build API url
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=imperial", city, apiKey)
	
	// Make the API call
	res, err := http.Get(url)
	if err != nil{
		return nil, err
	}
	defer res.Body.Close()

	// Read response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Parse JSON response
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}