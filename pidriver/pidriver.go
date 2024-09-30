package main

import (
    // "fmt"
    "log"
    "github.com/oppenheimer404/pidriver/pidriver/scanner"
)

const (
    REQUEST_WIFI_DATA = "wifi"
    REQUEST_BLUETOOTH_DATA = "bt"
)

func main() {
    response, err := scanner.Request(REQUEST_WIFI_DATA)
    if err != nil {
        log.Println(err)
    } else {
        log.Println(response)
    }
}
