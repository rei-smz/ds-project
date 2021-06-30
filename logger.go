package main

import (
	"fmt"
	"os"
	"time"
)

var log *os.File

func CreateLog() {
	var err error
	log, err = os.Create("log/" + time.Now().Format("20060102_150405") + ".log")
	if err != nil {
		fmt.Println("创建日志文件失败，退出程序")
		os.Exit(1)
	}
}

func Log(message string) {
	var err error
	message = "[" + time.Now().Format("2006-01-02 15:04:05") + "]" + message
	_, err = log.WriteString(message + "\n")
	if err != nil {
		fmt.Println( time.Now().Format("2006-01-02 15:04:05") + " 发生Log文件写入错误")
	}
}
