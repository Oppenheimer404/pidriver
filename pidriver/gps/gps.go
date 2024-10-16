package gps

import (
	"math/rand"
	"time"
)

// StartScan sends GPS data (in integer format) at the specified rate to a results channel.
func StartScan(rate time.Duration, results chan<- map[string]interface{}) {
	sendLocation(results)

	for range time.Tick(rate) {
		sendLocation(results)
	}
}

// sendLocation retrieves and sends the GPS location data to the results channel.
func sendLocation(results chan<- map[string]interface{}) {
	// Send example GPS data as integers
	results <- map[string]interface{}{
		"latitude":  randomInt(-90, 90),
		"longitude": randomInt(-180, 180),
		"altitude":  rand.Intn(1000), // Example altitude in meters
		"accuracy":  rand.Intn(10),   // Example accuracy in meters
	}
}

// randomInt generates a random integer in the specified range.
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
