package main

import (
	"fmt"
	"log"
    "flag"
	// "time"
	// "github.com/oppenheimer404/pidriver/pidriver/logger"
	"github.com/oppenheimer404/pidriver/pidriver/scanner"
	// "github.com/oppenheimer404/pidriver/pidriver/gps"
)

// Constant Variables
const (
    // scan requests
    REQUEST_WIFI_DATA = "wifi"
    REQUEST_BLUETOOTH_DATA = "bt"
    // logging setup
    FILETYPE_CSV = "csv"
)

func main() {
    // Ticker for testing purposes
    // ticker := time.NewTicker(1 * time.Second)
    // toggle := false
    // for range ticker.C {
    //     toggle = !toggle
    //     var data string
    //     if toggle {
    //         data = request(REQUEST_WIFI_DATA)
    //     } else if !toggle {
    //         data = request(REQUEST_BLUETOOTH_DATA)
    //     }
    //     fmt.Println(data)
    // }
}

func requestScan(requestType string) string {
    response, err := scanner.Request(requestType) // Collect data from scanner
    if err != nil {
        log.Fatal(err) // Fatal error if requestType is invalid
    }
    return response
}
