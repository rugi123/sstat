package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func getServerStats() []float64 {
	percent, _ := cpu.Percent(time.Second, false)
	mem, _ := mem.VirtualMemory()
	disk, _ := disk.Usage("/")
	stats := []float64{
		percent[0],
		mem.UsedPercent,
		float64(disk.Free) / 1024 / 1024 / 1024,
	}
	return stats
}
func main() {
	bot, err := tgbotapi.NewBotAPI("7473696296:AAEDG99W35oMBUV9htfP2P7ksd9Fwft28cc")
	if err != nil {
		fmt.Println("ошибка получения newbotapi:", err)
	}

	bot.Debug = true

	updconf := tgbotapi.NewUpdate(0)
	updconf.Timeout = 60

	updates := bot.GetUpdatesChan(updconf)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.Text == "stats" {
			if err != nil {
				fmt.Println("ошибка получения статов: ", err)
			}
			sstats := getServerStats()
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("🔥 CPU: %.1f%%\n🧩 MEM: %.1f%%\n💾 DISK: %.1fGB\n", sstats[0], sstats[1], sstats[2]))
			bot.Send(msg)
		}
	}
	/*
		fmt.Println(float64(disk.Free) / 1024 / 1024 / 1024)
		fmt.Println(mem.UsedPercent)
		fmt.Println(percent[0])
	*/
}
