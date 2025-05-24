package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
        "github.com/shirou/gopsutil/v3/mem"
        "github.com/shirou/gopsutil/v3/disk"
        "github.com/shirou/gopsutil/v3/sensors"
        "time"
)

func main() {
    percent, _ := cpu.Percent(time.Second, false)
    mem, _ := mem.VirtualMemory()
    disk, _ := disk.Usage("/")
    temp, _ := sensors.SensorsTemperatures()
    fmt.Println(temp[0])
    fmt.Println(float64(disk.Free)/1024/1024/1024)
    fmt.Println(mem.UsedPercent)
    fmt.Println(percent[0])
}
