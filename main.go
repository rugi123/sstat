package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	percent, _ := cpu.Percent(time.Second, false)
	mem, _ := mem.VirtualMemory()
	disk, _ := disk.Usage("/")
	fmt.Println(float64(disk.Free) / 1024 / 1024 / 1024)
	fmt.Println(mem.UsedPercent)
	fmt.Println(percent[0])
}
