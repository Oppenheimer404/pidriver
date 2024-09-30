package main

import (
    "fmt"
    "github.com/oppenheimer404/pidriver/pidriver/scanner"
)

func main() {
    x := scanner.Wifi()
    fmt.Println(x)
    x = scanner.BT()
    fmt.Println(x)
}
