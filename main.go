package main

import "sync"

var port int //服务端口
var wg sync.WaitGroup //用于多线程控制的等待组

func main() {
	//初始化部分
	CreateLog()
	Log("程序启动")
	status = Status{false, 0, []Navi {}}
	//结束初始化部分
	wg.Add(1)

	wg.Wait()
}
