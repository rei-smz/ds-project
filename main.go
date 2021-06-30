package main

import "sync"

var port int //服务端口
var wg sync.WaitGroup //用于多线程控制的等待组

func main() {
	//初始化部分
	logicalToPoint = make(map[string][]string)
	canteenCrowd = make(map[int]float64)
	locationInfo = make(map[string]string)
	pointIndex = make(map[Point]int)
	pointNameIndex = make(map[string]int)
	CreateLog()
	ReadConfig()
	CreateGraph()
	status = Status{false, 0, []Navi {}}
	Log("程序启动")
	//结束初始化部分
	wg.Add(1)
	go Server()
	wg.Wait()
}
