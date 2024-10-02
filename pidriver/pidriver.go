package main

import (
	"fmt"
	"log"
	"time"
	"your_project_path/wifi" // Import the wifi package
)

// Function that handles WiFi scanning and outputs results
func scanWiFi() {
	// Call the wifi.Request function to scan for networks
	networks, err := wifi.Request()
	if err != nil {
		log.Fatalf("failed to scan networks: %v", err)
	}

	// Print the details of each network
	fmt.Println("Visible WiFi Networks:")
	for _, network := range networks {
		fmt.Printf("SSID: %s, BSSID: %s, Frequency: %d MHz, Signal: %d dBm, Channel: %d\n",
			network.SSID, network.BSSID, network.Frequency, network.Signal, network.Channel)
	}
}

// rta function (same as before)
func rta(ch chan<- string) {
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for count := 0; count <= 100; count++ {
		<-ticker.C
		ch <- fmt.Sprintf("Routine A count: %d", count)
	}
	close(ch)
}

// rtb function (same as before)
func rtb(ch chan<- string) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for count := 0; count <= 100; count++ {
		<-ticker.C
		ch <- fmt.Sprintf("Routine B count: %d", count)
	}
	close(ch)
}

func main() {
	// Channels for rta and rtb
	rtaCh := make(chan string)
	rtbCh := make(chan string)

	// Start rta and rtb goroutines
	go rta(rtaCh)
	go rtb(rtbCh)

	// Combined channel for merged outputs
	combinedCh := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)

	// Forwarding channels to combinedCh
	go forward(rtaCh, combinedCh, &wg)
	go forward(rtbCh, combinedCh, &wg)

	go func() {
		wg.Wait()
		close(combinedCh)
	}()

	// Collect and print in batches of 10
	var results []string
	for msg := range combinedCh {
		results = append(results, msg)
		if len(results) == 10 {
			fmt.Println("Batch of 10 results:")
			for _, result := range results {
				fmt.Println(result)
			}
			fmt.Println("---------------------")
			results = nil
		}
	}

	// Print remaining results
	if len(results) > 0 {
		fmt.Println("Final batch of results:")
		for _, result := range results {
			fmt.Println(result)
		}
		fmt.Println("---------------------")
	}

	// WiFi scan after routines complete
	scanWiFi()

	fmt.Println("Main function completed")
}
