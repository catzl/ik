package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type Ret struct {
	Ret int    `json:"ret"`
	Msg string `json:"msg"`
}

type Config struct {
	Cookie   string `json:"cookie"`
	Pushplus string `json:"pushplus"`
}

func main() {
	iKuu()
}

func iKuu() {
	config := &Config{}
	osConfig := os.Getenv("CONFIG")
	json.Unmarshal([]byte(osConfig), config)

	url := "https://ikuuu.co/user/checkin"
	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Set("Cookie", config.Cookie)
	client := http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	ret := &Ret{}
	err := json.Unmarshal(body, ret)
	if err != nil {
		sendPushPlus(config.Pushplus, "iKuuu 签到", "Cookie 已失效请检查！")
		return
	}

	sendPushPlus(config.Pushplus, "iKuuu 签到", ret.Msg)
}

func sendPushPlus(token, title, content string) {
	url := "https://www.pushplus.plus/send"
	ma := make(map[string]interface{})
	ma["token"] = token
	ma["title"] = title
	ma["content"] = content
	js, _ := json.Marshal(ma)
	param := bytes.NewReader(js)

	req, _ := http.NewRequest("POST", url, param)
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
}
