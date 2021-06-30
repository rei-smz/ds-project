package main

import (
	"encoding/json"
	"fmt"
	"math"
	"time"
)

var pointList []Point //地点列表
var roadList []Road //道路列表
var logicalList []Logical
var status Status //状态信息
var canteenCrowd map[int]float64 //食堂拥挤度

func StartSimulation() {
	Log("开始模拟")
	status.IsRunning = true
	for _, navi := range status.NavigationList {
		if navi.Path[0].StartTime == -1 {
			navi.Path[0].StartTime = status.Time
		}
	}
	LogStatus()
	for status.IsRunning {
		time.Sleep(time.Second*1)
		status.Time += 5
		UpdateNavi()
		LogStatus()
	}
	wg.Done()
}

func UpdateNavi() {
	for i := 0; i < len(status.NavigationList); i ++ {
		if status.NavigationList[i].Distance <= 0 {
			continue
		}
		status.NavigationList[i].Time = status.Time - status.NavigationList[i].StartTime
		//status.NavigationList[i].RemainingTime -= status.NavigationList[i].Time
		var restTime = status.Time - status.NavigationList[i].Path[status.NavigationList[i].CurrentIndex].StartTime - status.NavigationList[i].Path[status.NavigationList[i].CurrentIndex].Time
		if status.Time - status.NavigationList[i].Path[status.NavigationList[i].CurrentIndex].StartTime >= status.NavigationList[i].Path[status.NavigationList[i].CurrentIndex].Time { //走完了某一路段
			status.NavigationList[i].Distance -= status.NavigationList[i].Path[status.NavigationList[i].CurrentIndex].Length
			if status.NavigationList[i].Distance <= 0 {
				Log(fmt.Sprintf("学生%s已到达目的地%s，导航结束。",
					status.NavigationList[i].StudentName,
					status.NavigationList[i].DestName))
				locationInfo[status.NavigationList[i].StudentName] = status.NavigationList[i].DestName
				continue
			}
			status.NavigationList[i].CurrentIndex++
			locationInfo[status.NavigationList[i].StudentName] = graph[status.NavigationList[i].Path[status.NavigationList[i].CurrentIndex].StartIndex].name
			Log(fmt.Sprintf("学生%s已到达%s，距离目的地还有%f米。",
				status.NavigationList[i].StudentName,
				locationInfo[status.NavigationList[i].StudentName],
				status.NavigationList[i].Distance))

			status.NavigationList[i].Path[status.NavigationList[i].CurrentIndex].StartTime = status.Time
			status.NavigationList[i].Path[status.NavigationList[i].CurrentIndex].Time -= restTime
		}
	}
	//for i := 0; i < len(status.NavigationList); i++ { //删除已结束的导航
	//	if status.NavigationList[i].Distance <= 0 {
	//		status.NavigationList[i].Path = []Trip{}
	//	}
	//}
}

func DelNavi(stuName string)  {
	for i := 0; i < len(status.NavigationList); i ++ {
		if status.NavigationList[i].StudentName == stuName {
			status.NavigationList = append(status.NavigationList[:i], status.NavigationList[i+1:]...)
			break
		}
	}
}

func LogStatus() {
	Log(fmt.Sprintf("模拟时间已过去%f秒，当前共有%d个模拟行程。",status.Time, len(status.NavigationList)))
}

func AddStudent(stuName string, location string) []byte {
	locationInfo[stuName] = location
	return []byte("success")
}

func AddNavi(student string,destination string, method int, mustPass []string) []byte {
	var navi Navi
	navi.StudentName = student
	for i := 0; i < len(status.NavigationList); i ++ {
		if status.NavigationList[i].StudentName == navi.StudentName {
			status.NavigationList = append(status.NavigationList[:i], status.NavigationList[i+1:]...)
			break
		}
	}
	//这里加上逻辑位置判断
	if _, ok := logicalToPoint[destination]; ok { //导航目的地为逻辑位置
		if destination == "食堂" {
			var minCrowd = math.Inf(1)
			var minIndex int
			for index, crowd := range canteenCrowd {
				if crowd < minCrowd && graph[index].schoolId == graph[pointNameIndex[locationInfo[student]]].schoolId {
					minIndex = index
					minCrowd = crowd
				}
			}
			navi.DestName = graph[minIndex].name
		} else {
			navi.DestName = logicalToPoint[destination][0]
		}
	} else {
		navi.DestName = destination
	}
	navi.StartTime = status.Time
	navi.Distance = 0
	navi.RemainingTime = 0
	navi.CurrentIndex = 0
	navi.Time = 0
	navi.Path = []Trip{}
	if method == 2 {
		for _, name := range mustPass {
			graph[pointNameIndex[name]].isNeeded = true
		}
		neededPath(pointNameIndex[locationInfo[student]], pointNameIndex[destination], len(mustPass), &navi)
	} else {
		findDis(pointNameIndex[locationInfo[student]], method)
		findPath(pointNameIndex[locationInfo[student]], pointNameIndex[destination], &navi, method)
	}
	status.NavigationList = append(status.NavigationList, navi)
	test, err := json.Marshal(status)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(test))
	}
	return []byte("success")
}

func GetNearby(stuName string) map[string][]nearbyInfo {
	var nearby []nearbyInfo
	var res map[string][]nearbyInfo
	res = make(map[string][]nearbyInfo)
	var start = pointNameIndex[locationInfo[stuName]]
	findDis(start, 0)
	for i := 0; i < len(graph); i ++ {
		if i == start {
			continue
		}
		if graph[i].dis <= 30 && graph[i].pointType == 0 {
			nearby = append(nearby, nearbyInfo{graph[i].name, graph[i].dis})
		}
	}
	res["nearby"] = nearby
	return res
}
