package gps

import (
	"fmt"
)

type RequestTypeGPS int

const (
	STATUS RequestTypeGPS = iota
	CURRENT
)

func Request(requestType RequestTypeGPS) (string, error) {
	switch requestType {
	case STATUS:
		return "Return Status", nil
	case CURRENT:
		return "Return Location", nil
	default:
		return "", fmt.Errorf("(gps.Request) invalid requestType: %d", requestType)
	}
}