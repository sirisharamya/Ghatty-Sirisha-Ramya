package main

import (
	"fmt"
	"strings"
)

type CityData struct {
	Name        string
	AverageTemp float64
	Rainfall    float64
}

func main() {
	// 1. Data Input: Hardcoded data for cities
	cities := []CityData{
		{"New York", 12.5, 1200},
		{"Los Angeles", 18.7, 380},
		{"Chicago", 10.3, 900},
		{"Houston", 22.1, 1400},
		{"Miami", 25.5, 1600},
	}

	// 2. Highest and Lowest Temperature
	highestTempCity, highestTemp := getCityWithHighestTemp(cities)
	lowestTempCity, lowestTemp := getCityWithLowestTemp(cities)

	fmt.Printf("City with the highest temperature: %s (%.2f°C)\n", highestTempCity, highestTemp)
	fmt.Printf("City with the lowest temperature: %s (%.2f°C)\n", lowestTempCity, lowestTemp)

	// 3. Average Rainfall
	averageRainfall := getAverageRainfall(cities)
	fmt.Printf("Average rainfall across all cities: %.2f mm\n", averageRainfall)

	// 4. Filter Cities by Rainfall
	var threshold float64
	fmt.Println("Enter the rainfall threshold (mm):")
	fmt.Scan(&threshold)
	filterCitiesByRainfall(cities, threshold)

	// 5. Search by City Name
	var cityName string
	fmt.Println("Enter city name to search for its data:")
	fmt.Scan(&cityName)
	searchCityByName(cities, cityName)
}

// Function to get the city with the highest temperature
func getCityWithHighestTemp(cities []CityData) (string, float64) {
	highestTemp := cities[0].AverageTemp
	highestTempCity := cities[0].Name

	for _, city := range cities {
		if city.AverageTemp > highestTemp {
			highestTemp = city.AverageTemp
			highestTempCity = city.Name
		}
	}
	return highestTempCity, highestTemp
}

// Function to get the city with the lowest temperature
func getCityWithLowestTemp(cities []CityData) (string, float64) {
	lowestTemp := cities[0].AverageTemp
	lowestTempCity := cities[0].Name

	for _, city := range cities {
		if city.AverageTemp < lowestTemp {
			lowestTemp = city.AverageTemp
			lowestTempCity = city.Name
		}
	}
	return lowestTempCity, lowestTemp
}

// Function to calculate the average rainfall across all cities
func getAverageRainfall(cities []CityData) float64 {
	var totalRainfall float64
	for _, city := range cities {
		totalRainfall += city.Rainfall
	}
	return totalRainfall / float64(len(cities))
}

// Function to filter cities by rainfall above a certain threshold
func filterCitiesByRainfall(cities []CityData, threshold float64) {
	fmt.Printf("Cities with rainfall above %.2f mm:\n", threshold)
	for _, city := range cities {
		if city.Rainfall > threshold {
			fmt.Printf("%s - %.2f mm\n", city.Name, city.Rainfall)
		}
	}
}

// Function to search for a city by name
func searchCityByName(cities []CityData, cityName string) {
	cityName = strings.Title(strings.ToLower(cityName))
	for _, city := range cities {
		if strings.Title(strings.ToLower(city.Name)) == cityName {
			fmt.Printf("Data for %s:\n", city.Name)
			fmt.Printf("Average Temperature: %.2f°C\n", city.AverageTemp)
			fmt.Printf("Rainfall: %.2f mm\n", city.Rainfall)
			return
		}
	}
	fmt.Println("City not found!")
}
