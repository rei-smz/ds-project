package main

import (
	"fmt"
	"time"
)

var pointList []Point //地点列表
var roadList []Road //道路列表
var status Status //状态信息
var schoolBus bool //是否有校车

func StartSimulation() {
	Log("开始模拟")
	status.isRunning = true
	schoolBus = false
	LogStatus()
	for status.isRunning {
		time.Sleep(time.Second*5) //每5秒钟对应模拟的1分钟
		status.time++
		if status.time % 10 == 0{
			schoolBus = true
		}else {
			schoolBus = false
		}
		UpdateNavi()
		LogStatus()
	}
	wg.Done()
}

func UpdateNavi() {
	for _, navi := range status.navigationList {
		navi.time = status.time - navi.startTime
		navi.remainingTime -= navi.time
		if navi.time - navi.path[navi.currentIndex].startTime == navi.path[navi.currentIndex].time { //走完了某一路段
			navi.distance -= navi.path[navi.currentIndex].length
			if navi.distance == 0 {
				Log(fmt.Sprintf("学生%s已到达目的地%s，导航结束。",
					navi.studentName,
					navi.destName))
				locationInfo[navi.studentName] = navi.destName
				continue
			}
			navi.currentIndex++
			Log(fmt.Sprintf("学生%s已到达%s，距离目的地还有%d米。",
				navi.studentName,
				graph[navi.path[navi.currentIndex].startIndex].name,
				navi.distance))
			navi.path[navi.currentIndex].startTime = status.time
		}
	}
	for i := 0; i < len(status.navigationList); i++ { //删除已结束的导航
		if status.navigationList[i].distance == 0 {
			status.navigationList = append(status.navigationList[:i], status.navigationList[i+1:]...)
			i--
		}
	}
}

func LogStatus() {
	Log(fmt.Sprintf("模拟时间已过去%d分钟，当前共有%d个模拟行程。",status.time, len(status.navigationList)))
}
