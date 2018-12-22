package util

import (
	"fmt"
)

// Pause 模拟暂停控制台操作
func Pause() {
	fmt.Print("按任意键继续......")
	fmt.Scanln()
}
