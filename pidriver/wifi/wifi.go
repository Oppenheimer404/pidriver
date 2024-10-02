package wifi

import (
	"github.com/mdlayher/wifi"
	"log"
)

// Reply represents the details of a WiFi network
type Reply struct {
	SSID          string // The SSID of the network
	BSSID         string // The MAC address of the access point
	Frequency     int    // Frequency in MHz
	Signal        int    // Signal strength in dBm
	Channel       int    // WiFi channel
}

// Request scans for visible WiFi networks and returns a list of Reply structs
func Request() ([]Reply, error) {
	// Open a new WiFi client using the nl80211 protocol
	client, err := wifi.New()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// Get list of available network interfaces
	ifis, err := client.Interfaces()
	if err != nil {
		return nil, err
	}

	// Scan for WiFi networks on the first available interface
	results, err := client.Scan(ifis[0])
	if err != nil {
		return nil, err
	}

	// Create a slice to hold the replies
	var replies []Reply

	// Loop over the results and create a Reply for each visible network
	for _, result := range results {
		replies = append(replies, Reply{
			SSID:      result.SSID,
			BSSID:     result.BSSID,
			Frequency: result.Frequency,
			Signal:    result.Signal,
			Channel:   result.Channel,
		})
	}

	return replies, nil
}
