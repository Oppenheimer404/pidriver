package scanner

import "fmt"

const (
	REQUEST_WIFI_DATA = "wifi"
	REQUEST_BLUETOOTH_DATA = "bt"
)

func Request(requestType string) (string, error) {
	switch requestType {
	case REQUEST_WIFI_DATA:
		return scanWifi(), nil
	case REQUEST_BLUETOOTH_DATA:
		return scanBT(), nil
	default:
		return "", fmt.Errorf("invalid request type: %s", requestType)
	}
}

func scanWifi() string {
	x := "Wifi Data"
	return x
}

func scanBT() string {
	x := "Bluetooth Data"
	return x
}