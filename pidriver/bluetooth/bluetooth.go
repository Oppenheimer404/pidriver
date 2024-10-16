package bluetooth

import (
	"crypto/rand"
	"fmt"
	r "math/rand"
	"time"
)

// StartScan sends Bluetooth data (with mixed types) at the specified rate to a results channel.
func StartScan(rate time.Duration, results chan<- map[string]interface{}) {
	// Send initial device info
	sendDeviceInfo(results)

	// Continuously send device info based on the specified rate
	for range time.Tick(rate) {
		sendDeviceInfo(results)
	}
}

// sendDeviceInfo sends simulated Bluetooth network information as a map with mixed types
func sendDeviceInfo(results chan<- map[string]interface{}) {
	results <- map[string]interface{}{
		"CellKey":            randomCellKey(),                 // Example Cell Key as a string
		"NetworkName":        "Foobar",                        // Example Network Name as a string
		"Capabilities":       "Uncategorized;10",              // Capabilities as string
		"FirstTimestampSeen": time.Now().Format(time.RFC3339), // Timestamp as string
		"Channel":            randomInt(1, 11),                // Bluetooth channel as int
		"Frequency":          randomInt(2400, 5800),           // Bluetooth frequency as int
		"RSSI":               randomInt(-100, -30),            // Signal strength as int
		"Type":               "BT",                            // Device type as string
	}
}

// randomInt generates a random integer between min and max.
func randomInt(min, max int) int {
	return min + r.Intn(max-min)
}

// randomCellKey generates a random Cell Key string.
func randomCellKey() string {
	bssid := make([]byte, 6)
	rand.Read(bssid)
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", bssid[0], bssid[1], bssid[2], bssid[3], bssid[4], bssid[5])
}
