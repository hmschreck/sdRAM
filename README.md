# RAM utilization plugin for Elgato Stream Deck
A simple RAM utilization plugin for the Elgato Stream Deck keypad.

# Compiling
## Prerequisites
* Golang 1.8+
* github.com/shirou/gopsutil
* github.com/StackOverflow/wmi (on Windows)

`go build -o mem.exe main.go` on Windows

`go build -o mem main.go` on Mac