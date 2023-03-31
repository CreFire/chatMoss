package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"net/url"
)

func main() {

	botToken := viper.GetString("telegramApiToken")
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
	bot, err := tgbotapi.NewBotAPIWithClient(botToken, tgbotapi.APIEndpoint, client)
	if err != nil {
		panic(any(err))
	}
	bot.Debug = true
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
