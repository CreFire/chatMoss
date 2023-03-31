package main

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	ur := "https://api.chat.com/v1/messages"
	method := "POST"

	payload := strings.NewReader(`{
    "recipient": "user@example.com",
    "message": "Hello, World!"
}`)
	port := viper.Get("proxy.port")
	uri := fmt.Sprintf("http://127.0.0.1:%d", port)
	logrus.Info("uri:%v", uri)
	u, _ := url.Parse(uri)
	client := &http.Client{
		Transport: &http.Transport{
			// 设置代理
			Proxy: http.ProxyURL(u),
		},
	}
	req, err := http.NewRequest(method, ur, payload)

	if err != nil {
		logrus.Println(err)
		return
	}

	req.Header.Add("Content-Type", "application/json")
	token := fmt.Sprintf("Bearer %v", viper.Get("api.chatgptfourapi"))
	logrus.Println("token:", token)
	req.Header.Add("Authorization", token)
	res, err := client.Do(req)
	if err != nil {
		logrus.Errorf("err:%v", err)
		return
	}
	defer res.Body.Close()

	if err != nil {
		logrus.Println(err)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logrus.Println(err)
		return
	}

	var response map[string]interface{}
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		logrus.Error("err:%v", err)
		return
	}

	logrus.Println(response)

}

func init() {
	viper.SetConfigFile("./config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Errorf("config err:%v", err)
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file:%s Op:%s\n", e.Name, e.Op)
	})
	fmt.Printf("config: %v \n %v", viper.AllKeys(), viper.AllSettings())
}
