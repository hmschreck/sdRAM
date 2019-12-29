package main

import (
	"fmt"
	"github.com/valyala/fastjson"
	"meow.tf/streamdeck/sdk"
	"os"
	"time"
	"github.com/shirou/gopsutil/mem"
)

var ticker = time.NewTicker(1 * time.Second)
var done = make(chan bool)
var debug = false
var log, _ = os.Create("log")
func main() {
	defer log.Close()
	log.WriteString(fmt.Sprintf("%v", os.Args))
	if debug {log.WriteString("In Main\n")}
	sdk.RegisterAction("com.hmschreck.memory.pct", cpuTempHandler)
	err := sdk.Open()
	if err != nil {
		sdk.Log("Died")
	}
	if debug {log.WriteString("Successfully opened up the connection to the SDK\n")}
	sdk.Wait()
}

func cpuTempHandler(action, context string, payload *fastjson.Value, deviceId string) {
	if debug {
		log.WriteString(action)
		log.WriteString("\n")
		log.WriteString(context)
		log.WriteString("\n")
		log.WriteString(payload.String())
		log.WriteString("\n")
	}
	sdk.Log(action)
	sdk.Log(context)
	sdk.Log(fmt.Sprintf("%v", payload))
	if action == "com.hmschreck.memory.pct" {
		go func() {
			for {
				select {
				case <-done:
					return
				case _ = <-ticker.C:
					v, _ := mem.VirtualMemory()
					sdk.SetTitle(context, fmt.Sprintf("RAM\n%.0f%%", v.UsedPercent ), sdk.TargetBoth)
				}
			}
		}()
	}
	if action == "willDisappear" {
		done <- true
	}
}
