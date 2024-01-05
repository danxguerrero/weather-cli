package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "io"
    "net/http"
)

const backendURL = "https://weather-cli-backend.onrender.com/weather" 

func getWeather(city string) (map[string]interface{}, error) {
    url := fmt.Sprintf("%s?city=%s", backendURL, city)

    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var weatherData map[string]interface{}
    err = json.Unmarshal(body, &weatherData)
    if err != nil {
        return nil, err
    }

    return weatherData, nil
}

func main() {
    var city string
    flag.StringVar(&city, "city", "", "City name for which to fetch weather data")
    flag.Parse()

    if city == "" {
        fmt.Println("Please provide a city name using the -city flag")
        return
    }

    weatherData, err := getWeather(city)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("---", city, "---")
    fmt.Printf("Current temperature: %v°F\n", weatherData["main"].(map[string]interface{})["temp"])
}
