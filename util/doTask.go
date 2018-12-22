package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

//DoTask 模拟开红包
func DoTask(phone, redpackID string) {
	fmt.Println("**************************************" + phone + "*****************************************")
	url := "https://rb.bitkeep.org/reciveBag"

	payload := strings.NewReader("rid=" + redpackID + "&username=" + phone)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cache-Control", "no-cache")

	res, _ := http.DefaultClient.Do(req)
	if res != nil {
		defer res.Body.Close()
	}
	body, _ := ioutil.ReadAll(res.Body)
	respStr := string(body)
	errnoObj := gjson.Get(respStr, "errno")
	if errnoObj.Exists() {
		code := errnoObj.Int()
		if code == 0 {
			fmt.Printf("username:%s, 开红包成功\n", phone)
		} else {
			msg := gjson.Get(respStr, "msg").String()
			fmt.Printf("username:%s, 开红包失败, %s\n", phone, msg)
		}
	} else {
		fmt.Printf("username:%s, 开红包失败, errnoObj不存在\n", phone)
	}
}
