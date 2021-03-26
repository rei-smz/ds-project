package main

/*type Location struct {
	name string
	locationName string
}*/

var locationInfo map[string]string //学生位置信息，格式为["名字"]: "位置"

type Point struct {
	pointType 	int //地点类型，0为设施或楼层，1为路口或校门等其他地点
	name 		string
}

type Road struct {
	startName 	   string
	endName        string
	length         int
	isBikeRoad     bool
	crowdedForWalk float64 //步行的拥挤度
	crowdedForBike float64 //骑车的拥挤度
}

type Arc struct {
	targetIndex    int
	length         int
	isBikeRoad     bool
	crowdedForWalk float64 //步行的拥挤度
	crowdedForBike float64 //骑车的拥挤度
}

type Vertex struct {
	name 	string
	arcList []Arc
}

type Trip struct {
	startIndex 	int //当前行程路段的起点（使用图的顶点编号）
	arcIndex 	int //当前行程路段走的边（使用起点边列表中的编号）
	time 		int //已用时间
}

type Navi struct {
	studentName 	string
	distance 		int //目的地距离
	time 			int //已用时间
	remainingTime 	int //剩余时间
	currentIndex 	int //当前走到的
	path 			[]Trip //导航路径
}

type Status struct {
	isRunning 		bool
	time 			int
	navigationList []Navi
}



var logicalToPoint map[string]Point //逻辑位置与地点的映射
var pointIndex map[Point]int //地点与图中索引值的映射
