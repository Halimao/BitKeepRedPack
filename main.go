package main

import (
	"BitKeepRedPack/util"
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"unicode/utf8"
)

// wg 同步锁
var wg sync.WaitGroup

func getRedpackID(redpackURL string) string {
	length := utf8.RuneCountInString(redpackURL)
	index := strings.LastIndex(redpackURL, "/") + 1
	return redpackURL[index:length]
}

// tryCatch 实现 try catch.
// 空接口：具有0个方法的接口称为空接口。它表示为interface {}。由于空接口有0个方法，所有类型都实现了空接口
func tryCatch(try func(phone, redpackID string), handler func(interface{}), phone string, redpackID string) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
		wg.Done()
	}()
	try(phone, redpackID)
}

func main() {
	phoneArray := util.InitConfig()
	redpackURL := ""
	for {
		fmt.Println("Please input redpackURL(Type 'quit' to leave): ")
		fmt.Scanln(&redpackURL)
		redpackID := getRedpackID(redpackURL)
		if redpackID == "quit" {
			break
		}
		fmt.Println("redpackID--", redpackID)
		if len(phoneArray) > 0 {
			for _, value := range phoneArray {
				wg.Add(1)
				go tryCatch(util.DoTask, func(err interface{}) { fmt.Printf("phone:%s,抢红包异常, %v\n", value, err) }, value, redpackID)
			}
		}
		wg.Wait()
		cmd := exec.Command("cmd", "/c start https://rb.bitkeep.org/blist/" + redpackID)
		if err := cmd.Start(); err != nil {
			fmt.Println(err)
		}
	}
	util.Pause()
}
