package main

import (
	"fmt"
	"testing"
	"time"
)

// Login 模拟登录
func TestAny(t *testing.T) {
	fmt.Println(time.Now().Day())
	fmt.Println(time.Now().AddDate(0, 0, -1).Day())
}
