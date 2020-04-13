package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strings"
)

type LogData struct {
	Topic string `json:"topic"`
	Data string `json:"data"`
}

var (
	client *elastic.Client
	esCh = make(chan *LogData, 10000)
)

// 初始化es，准备接收kafka那边发过来的数据
func Init(addr string) (err error){
	if !strings.HasPrefix(addr,"http") {
		addr = "http://"+addr
	}
	client, err = elastic.NewClient(elastic.SetURL(addr))
	if err != nil {
		panic(err)
		return
	}
	go sendToES()
	fmt.Println("connect to es success")
	return
}

// 发送数据到es
func SendToESChan(msg *LogData) {
	esCh <- msg
}

// 从通道中取数据发送到es
func sendToES() {
	for {
		select {
		case msg :=<- esCh:
			put1, err := client.Index().Index(msg.Topic).BodyJson(msg.Data).Do(context.Background())
			if err != nil {
				// Handler error
				fmt.Println(err)
			}
			fmt.Printf("Indexed success %v to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
		}
	}
}