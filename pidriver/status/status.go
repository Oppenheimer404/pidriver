package status

import "fmt"

type Request int

const (
	ALL Request = iota
	GPS
	WIFI
	BLUETOOTH
)

func Verify(device Request) (bool, error) {
	switch device {
	case ALL:
		return true, nil
	case GPS:
		return false, fmt.Errorf("(status.Verify) wip")
	case WIFI:
		return false, fmt.Errorf("(status.Verify) wip")
	case BLUETOOTH:
		return false, fmt.Errorf("(status.Verify) wip")
	default:
		return false, fmt.Errorf("(status.Verify) improper device specified: %d", device)
	}
}