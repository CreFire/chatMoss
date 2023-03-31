package main

import (
	"github.com/chatMoss/internal"
	"github.com/chatMoss/model"
	"log"
	"net/http"
	"sync"
)

func main() {
	chat := &model.Chat{
		Messages: []string{},
	}
	r := internal.Route()

	r.LoadHTMLFiles("index.html")
	log.Println("Starting server on port 8080...")
	var gw sync.WaitGroup
	gw.Add(1)
	go func() {
		defer gw.Done()
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
	gw.Wait()
}
