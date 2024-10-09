package wifi





// package wifi

// import (
//     "fmt"
//     "time"
// )

// type ScanType int

// const (
//     ALL_DATA ScanType = iota
// )

// // startScan simulates a process that runs at a specified rate for a given duration.
// func startScan(ch chan<- string, rate int, duration time.Duration) {
// 	ticker := time.NewTicker(time.Duration(1000/rate) * time.Millisecond) // Ticker based on rate (times per second)
// 	defer ticker.Stop()

// 	timer := time.NewTimer(duration) // Timer to stop after duration
// 	defer timer.Stop()

// 	count := 0

// 	// Loop until the timer expires
// 	for {
// 		select {
// 		case <-ticker.C:
// 			count++
// 			ch <- fmt.Sprintf("tick %d", count) // Send count to the channel

// 		case <-timer.C:
// 			close(ch) // Stop when the duration has passed
// 			return
// 		}
// 	}
// }

// func packageData() {
//     scanData := make(chan string)

// 	// Start the scan in a goroutine
// 	go startScan(scanData, 3, 1*time.Minute)

// 	var results []string
// 	for msg := range scanData {
// 		results = append(results, msg)
// 		if len(results) == 10 {
// 			// Print the current batch of 10 results
// 			fmt.Println("Batch of 10 results:")
// 			for _, result := range results {
// 				fmt.Println(result)
// 			}
// 			fmt.Println("---------------------")
// 			results = nil
// 		}
// 	}

// 	// Print any remaining results after the loop
// 	if len(results) > 0 {
// 		fmt.Println("Final batch of results:")
// 		for _, result := range results {
// 			fmt.Println(result)
// 		}
// 		fmt.Println("---------------------")
// 	}

// 	fmt.Println("Main function completed")
// }

// func Request() {
    
//     }