package wifi

import (
	"crypto/rand"
	"fmt"
	r "math/rand"
	"time"
)

// StartScan sends WiFi data (with mixed types) at the specified rate to a results channel.
func StartScan(rate time.Duration, results chan<- map[string]interface{}) {
	// Send initial device info
	sendDeviceInfo(results)

	// Continuously send device info based on the specified rate
	for range time.Tick(rate) {
		sendDeviceInfo(results)
	}
}

const (
	BSSID = "BSSID"
	SSID = "SSID"
	CAPABILITIES = "Capabilities"
	FIRST_TIME_SEEN = "FirstTimestampSeen"
	CHANNEL = "Channel"
	FREQUENCY = "Frequency"
	RSSI = "RSSI"
	TYPE = "Type"
)

// sendDeviceInfo sends simulated WiFi network information as a map with mixed types
func sendDeviceInfo(results chan<- map[string]interface{}) {
	results <- map[string]interface{}{
		BSSID:              randomBSSID(),                   // Example BSSID as a string
		SSID:               "Foobar",                        // Example SSID as a string
		CAPABILITIES:       "[WPA2]",                        // Capabilities as string
		FIRST_TIME_SEEN: time.Now().Format(time.RFC3339), // Timestamp as string
		CHANNEL:            randomInt(1, 11),                // WiFi channel as int
		FREQUENCY:          randomInt(2400, 5800),           // WiFi frequency as int
		RSSI:               randomInt(-100, -30),            // Signal strength as int
		TYPE:               "WIFI",                          // Device type as string
	}
}

// randomInt generates a random integer between min and max.
func randomInt(min, max int) int {
	return min + r.Intn(max-min)
}

// randomBSSID generates a random BSSID string.
func randomBSSID() string {
	bssid := make([]byte, 6)
	rand.Read(bssid)
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", bssid[0], bssid[1], bssid[2], bssid[3], bssid[4], bssid[5])
}
