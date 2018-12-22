package util

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tidwall/gjson"
)

// InitConfig 初始化配置
func InitConfig() []string {
	file, err := os.Open("conf/conf.json")
	defer file.Close()
	var phones []string
	if err != nil {
		fmt.Println("Initialize config failed, cause:", err)
		return phones
	}
	conf, _ := ioutil.ReadAll(file)
	confStr := string(conf)
	phonesObj := gjson.Get(confStr, "phones")
	if phonesObj.Exists() {
		for _, value := range phonesObj.Array() {
			phone := value.Get("phone").String()
			phones = append(phones, phone)
		}
	}
	fmt.Println("Initialize config success")
	return phones
}
