package main

import (
	"fmt"
	"os"
	"time"
)

var log *os.File

func CreateLog() {
	var err error
	log, err = os.Create("log/" + time.Now().Format("20210326_171046") + ".log")
	if err != nil {
		fmt.Println("创建日志文件失败，退出程序")
		os.Exit(1)
	}
}

func Log(message string) {
	var err error
	message = "[" + time.Now().Format("2021-3-26 17:15:19") + "]" + message
	_, err = log.WriteString(message + "\n")
	if err != nil {
		fmt.Println( time.Now().Format("2021-3-26 17:15:19") + " 发生Log文件写入错误")
	}
}
