package main

import (
	"math"
	"math/rand"
	"strings"
	"time"
)


var schoolBusRunTime = 1.0 //校车运行需要的时间，设置成比公共交通要短的时间

func CreateGraph()  {
	for i, p := range pointList {
		var newVertex Vertex
		newVertex.name = p.PointName
		newVertex.pointType = p.PointType
		newVertex.number = 0
		newVertex.schoolId = p.schoolId
		newVertex.pre = -1
		newVertex.preArc = -1
		newVertex.book = false
		newVertex.dis = 0
		newVertex.isNeeded = false
		newVertex.arcList = []Arc{}
		pointNameIndex[p.PointName] = i
		if strings.Contains(p.PointName, "食堂") {
			rand.Seed(time.Now().Unix())
			canteenCrowd[i] = rand.Float64()
		}
		graph = append(graph, newVertex)
	}
	for _, a := range roadList {
		var newArc1, newArc2 Arc
		newArc1.targetIndex = pointNameIndex[a.EndName]
		newArc2.targetIndex = pointNameIndex[a.StartName]
		newArc1.length = a.Length
		newArc2.length = a.Length
		newArc1.isBikeRoad = a.IsBikeRoad
		newArc2.isBikeRoad = a.IsBikeRoad
		rand.Seed(time.Now().Unix())
		newArc1.crowdedForBike = rand.Float32()
		if newArc1.crowdedForBike == 0 {
			newArc1.crowdedForBike += 0.001
		}
		newArc1.crowdedForWalk = rand.Float32()
		if newArc1.crowdedForWalk == 0 {
			newArc1.crowdedForWalk += 0.001
		}
		newArc2.crowdedForBike = newArc1.crowdedForBike
		newArc2.crowdedForWalk = newArc1.crowdedForWalk
		if graph[pointNameIndex[a.StartName]].schoolId != graph[pointNameIndex[a.EndName]].schoolId {
			newArc1.timeWalk = newArc1.length / 15
			newArc1.timeBike = newArc1.length / 15
			newArc2.timeWalk = newArc1.timeWalk
			newArc2.timeBike = newArc1.timeBike
		} else {
			newArc1.timeWalk = newArc1.length / (1.5 * float64(newArc1.crowdedForWalk))
			newArc1.timeBike = newArc1.length / (10 * float64(newArc1.crowdedForBike))
			newArc2.timeWalk = newArc1.timeWalk
			newArc2.timeBike = newArc1.timeBike
		}
		graph[pointNameIndex[a.StartName]].arcList = append(graph[pointNameIndex[a.StartName]].arcList, newArc1)
		graph[pointNameIndex[a.StartName]].number ++
		graph[pointNameIndex[a.EndName]].arcList = append(graph[pointNameIndex[a.EndName]].arcList, newArc2)
		graph[pointNameIndex[a.EndName]].number ++
	}
	Log("地图创建成功")
}

func NaviInit(start int, flag int)  {
	for i := 0; i < len(graph); i++ {
		graph[i].dis = math.Inf(1) //正无穷
		graph[i].book = false
		graph[i].pre = i
		graph[i].preArc = 0
	}
	graph[start].dis = 0
	graph[start].book = true
	for i := 0; i < graph[start].number; i++ {
		var newDis float64
		if flag == 0 {
			newDis = graph[start].arcList[i].length
		} else if flag == 1 {
			newDis = graph[start].arcList[i].timeWalk
		} else if flag == 3 {
			if !graph[start].arcList[i].isBikeRoad {
				newDis = graph[start].arcList[i].timeWalk
			} else {
				newDis = graph[start].arcList[i].timeBike
			}
		}
		endIndex := graph[start].arcList[i].targetIndex
		graph[endIndex].dis = newDis
		graph[endIndex].pre = start
		graph[endIndex].preArc = i
	}
}

func findDis(start int, flag int)  {
	NaviInit(start, flag)
	for i := 0; i < len(graph); i++ {
		minNum := math.Inf(1)
		indexNum := 0
		for j := 0; j < len(graph); j++ {
			if graph[j].dis < minNum && graph[j].book == false {
				minNum = graph[j].dis
				indexNum = j
			}
			/*if graph[j].book == false {
				indexNum = j
				break
			}*/
		}
		graph[indexNum].book = true

		for j := 0; j <graph[indexNum].number; j++ {
			endIndex := graph[indexNum].arcList[j].targetIndex
			var u float64
			if flag == 0 {
				u = graph[indexNum].arcList[j].length
			} else if flag == 1 {
				u = graph[indexNum].arcList[j].timeWalk
			} else if flag == 3 {
				if !graph[indexNum].arcList[j].isBikeRoad {
					u = graph[indexNum].arcList[j].timeWalk
				} else {
					u = graph[indexNum].arcList[j].timeBike
				}
			}
			if graph[endIndex].dis > graph[indexNum].dis + u {
				graph[endIndex].dis = graph[indexNum].dis + u
				graph[endIndex].pre = indexNum
				graph[endIndex].preArc = j
			}
		}
	}
}

func findPath(start int, end int, naviEdit *Navi, method int)  {
	kk := end
	var path Stack
	var arcNumber Stack
	path.Push(kk)
	arcNumber.Push(graph[kk].preArc)
	for kk != start {
		kk = graph[kk].pre
		path.Push(kk)
		if kk != start {
			arcNumber.Push(graph[kk].preArc)
		}
	}
	for !path.isEmpty() {
		point, _ := path.Top()
		if !arcNumber.isEmpty() {
			number, _ := arcNumber.Top()
			var trip Trip
			trip.StartIndex = point.(int)
			trip.ArcIndex = number.(int)
			trip.Length = graph[point.(int)].arcList[number.(int)].length
			if graph[point.(int)].schoolId != graph[graph[point.(int)].arcList[number.(int)].targetIndex].schoolId {
				if (naviEdit.RemainingTime - naviEdit.StartTime - (float64(int64(naviEdit.RemainingTime - naviEdit.StartTime) / 10) * 10) + schoolBusRunTime) < graph[point.(int)].arcList[number.(int)].timeWalk {
					trip.Time = naviEdit.RemainingTime - naviEdit.StartTime - (float64(int64(naviEdit.RemainingTime - naviEdit.StartTime) / 10) * 10) + schoolBusRunTime
				} else {
					trip.Time = graph[point.(int)].arcList[number.(int)].timeWalk // 跨校区不论导航策略都将公共交通时间设为timeWalk
				}
			} else {
				if method == 3 {
					trip.Time = graph[point.(int)].arcList[number.(int)].timeBike
				} else {
					trip.Time = graph[point.(int)].arcList[number.(int)].timeWalk
				}
			}
			trip.StartTime = -1
			trip.Passed = 0
			naviEdit.Path = append(naviEdit.Path, trip)
			naviEdit.RemainingTime += trip.Time
			naviEdit.Distance += trip.Length
			arcNumber.Pop()
		}
		path.Pop()
	}
}

func neededPath(start int, end int, num int, naviEdit *Navi)  {
	if num == 0 {
		findDis(start,1)
		findPath(start, end, naviEdit, 1)
		return
	}
	findDis(start, 1)
	minNum := math.Inf(1)
	index := 0
	for i := 0; i < len(graph); i++ {
		if i == start {
			continue
		}
		if graph[i].isNeeded == true {
			if graph[i].dis < minNum {
				minNum = graph[i].dis
				index = i
			}
		}
	}
	graph[index].isNeeded = false
	findPath(start, index, naviEdit, 1)
	neededPath(index, end, num - 1,naviEdit)
}
