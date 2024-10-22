// Logging package will not throw fatal errors
package logging

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/oppenheimer404/pidriver/pidriver/bluetooth"
	"github.com/oppenheimer404/pidriver/pidriver/gps"
	"github.com/oppenheimer404/pidriver/pidriver/wifi"
)

const (
	GET_TYPE      = "Type"
	DATATYPE_WIFI = "WIFI"
	DATATYPE_BT   = "BT"

	// strings for fixedData map
	A = "BSSID or CellKey"
	B = "SSID or NetworkName"
	C = "Capabilities"
	D = "FirstTimestampSeen"
	E = "Channel"
	F = "Frequency"
	G = "RSSI"
	H = "Latitude"
	I = "Longitude"
	J = "Altitude"
	K = "Accuracy"
	L = "RCOIs"
	M = "Mfgrid"
	N = "Type"
)

func mendWifi(deviceData map[string]interface{}, gpsData map[string]interface{}) map[string]interface{} {
	fixedData := make(map[string]interface{})

	fixedData[A] = deviceData[wifi.BSSID]
	fixedData[B] = deviceData[wifi.SSID]
	fixedData[C] = deviceData[wifi.CAPABILITIES]
	fixedData[D] = deviceData[wifi.CHANNEL]
	fixedData[E] = deviceData[wifi.FREQUENCY]
	fixedData[F] = deviceData[wifi.RSSI]
	fixedData[G] = gpsData[gps.LAT]
	fixedData[H] = gpsData[gps.LON]
	fixedData[I] = gpsData[gps.ALT]
	fixedData[K] = gpsData[gps.ACCURACY]
	fixedData[L] = ""
	fixedData[M] = ""
	fixedData[N] = "WIFI"

	return fixedData
}

func mendBluetooth(deviceData map[string]interface{}, gpsData map[string]interface{}) map[string]interface{} {
	fixedData := make(map[string]interface{})

	fixedData[A] = deviceData[bluetooth.CELL_KEY]
	fixedData[B] = deviceData[bluetooth.NETWORK_NAME]
	fixedData[C] = deviceData[bluetooth.CAPABILITIES]
	fixedData[D] = deviceData[bluetooth.CHANNEL]
	fixedData[E] = deviceData[bluetooth.FREQUENCY]
	fixedData[F] = deviceData[bluetooth.RSSI]
	fixedData[G] = gpsData[gps.LAT]
	fixedData[H] = gpsData[gps.LON]
	fixedData[I] = gpsData[gps.ALT]
	fixedData[K] = gpsData[gps.ACCURACY]
	fixedData[L] = ""
	fixedData[M] = ""
	fixedData[N] = "BT"

	return fixedData
}

func appendCSV(filename string, loggingData map[string]interface{}) error {
	// Open the CSV file in append mode
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close() // Ensure the file is closed after we're done

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush() // Flush any buffered data

	// Define the order of keys from A to N
	keys := []string{A, B, C, D, E, F, G, H, I, J, K, L, M, N}

	// Convert the map to a slice of strings in the specified order
	var record []string
	for _, key := range keys {
		value, exists := loggingData[key]
		if !exists {
			record = append(record, "") // Append an empty string if the key doesn't exist
			continue
		}
		// Convert value to string and append to the record slice
		switch v := value.(type) {
		case string:
			record = append(record, v)
		case int:
			record = append(record, strconv.Itoa(v)) // Convert int to string
		case float64:
			record = append(record, strconv.FormatFloat(v, 'f', -1, 64)) // Convert float64 to string
		default:
			record = append(record, fmt.Sprintf("%v", v)) // Fallback for other types
		}
	}

	// Write the logging data as a new record
	if err := writer.Write(record); err != nil {
		return fmt.Errorf("error writing record to file: %w", err)
	}

	return nil // Return nil if successful
}

func Default(deviceData map[string]interface{}, gpsData map[string]interface{}) {
	// Set dataType from device data
	dataType := deviceData[GET_TYPE].(string)
	loggingData := make(map[string]interface{})
	switch dataType {
	case DATATYPE_WIFI:
		loggingData = mendWifi(deviceData, gpsData)
	case DATATYPE_BT:
		loggingData = mendBluetooth(deviceData, gpsData)
	default:
		fmt.Println("Unknown data type:", dataType)
	}

	if err := appendCSV("data.csv", loggingData); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Data appended successfully!")
	}
}

func Error(err error, info string) {
	fmt.Printf("Error Logged: [%v] Reason: [%s] \n", err, info)
}
