package diagnose

import (
	"github.com/oppenheimer404/pidriver/pidriver/logging"
)

func Error(err error) {

	// Returns immediately if no error ocurred
	if err == nil {
		return
	}

	// Convert potential error to string
	errorString := err.Error()

	// Ensure the error string has at least one character to avoid panic
	if len(errorString) < 1 {
		logging.Error("[Diagnose Error] error string is empty")
		return
	}

	severity := string(errorString[0])

	switch severity {
	case "0": // Fatal Error
		return
	case "1": // Warning
		return
	case "2": // Info
		return
	default:
		logging.Error("[Diagnose Error] no severity indicated")
		logging.Error(string(errorString[1:]))
	}
}
