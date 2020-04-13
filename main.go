package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"llvvlv00.org/logtransfer/conf"
	"llvvlv00.org/logtransfer/es"
	"llvvlv00.org/logtransfer/kafka"
)

// log transfer
// 将日志数据从kafka取出来发往Es

var cfg conf.LogTransferCfg

func main() {
	//1、加载配置文件
	if err := ini.MapTo(&cfg, "./conf/config.ini"); err != nil {
		fmt.Println("加载配置文件失败 ", err)
		return
	}
	fmt.Println("加载配置文件成功!", cfg)

	//2、初始化
		//2、1 初始化一个es链接的client
	if err := es.Init(cfg.ESCfg.Address);err!= nil{
		fmt.Printf("init ES client failed, err:%v\n", err)
		return
	}

	//2、2初始化kafka链接的client
	if err := kafka.Init([]string{cfg.KafkaCfg.Address}, cfg.KafkaCfg.Topic); err != nil {
		fmt.Println("初始化 kafka 失败 ", err)
		return
	}

	//3、 从kafka取日志数据
	select {

	}

	//4、发往ES
}
