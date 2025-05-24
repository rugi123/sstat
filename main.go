package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command(`top -bn1 | awk '{printf "CPU: %.1f%%\n", $2 + $4}'`)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println("Ошибка выполнения команды:", err)
		return
	}

	fmt.Println("Результат выполнения команды:")
	fmt.Println(out.String())
}
