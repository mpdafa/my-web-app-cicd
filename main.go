package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	http.HandleFunc("/", sysInfoHandler)
	fmt.Println("Server starting on port 8888...")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func sysInfoHandler(w http.ResponseWriter, r *http.Request) {
	cpuUsages, _ := cpu.Percent(time.Second, true)
	memStat, _ := mem.VirtualMemory()
	diskStat, _ := disk.Usage("/")

	// Create a response string
	response := "System Information\n"
	response += "CPU Usage:\n"
	for i, usage := range cpuUsages {
		response += fmt.Sprintf(" CPU %d: %.2f%%\n", i, usage)
	}

	response += fmt.Sprintf("\nMemory Usage:\n Total: %v MB, Available: %v MB, Used: %.2f%%\n",
		memStat.Total/1024/1024, memStat.Available/1024/1024, memStat.UsedPercent)

	response += fmt.Sprintf("\nDisk Usage (for /):\n Total: %v GB, Free: %v GB, Used: %.2f%%\n",
		diskStat.Total/1024/1024/1024, diskStat.Free/1024/1024/1024, diskStat.UsedPercent)

	// Write the response to the HTTP ResponseWriter
	fmt.Fprint(w, response)
}
