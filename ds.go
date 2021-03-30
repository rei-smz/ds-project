package main


var locationInfo map[string]string //学生位置信息，格式为["名字"]: "位置"

type Point struct {
	pointType 	int //地点类型，0为设施或楼层，1为路口或校门等其他地点
	schoolId 	int //校区编号
	name 		string
}

type Road struct {
	startName 	   string
	endName        string
	schoolId 	   int //校区编号
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
	length 		int //路段长度，规划路径时赋值
	passed 		int //路段上的已走距离
	startTime 	int
	time 		int //需要时间
}

type Navi struct {
	studentName 	string
	distance 		int //目的地距离
	destName 		string //目的地名称
	startTime 		int //开始时间，添加导航时设为status.time
	time 			int //已用时间，每次更新导航状态时赋值为status.time-startTime
	remainingTime 	int //剩余时间
	currentIndex 	int //当前走到的path中的路段编号
	path 			[]Trip //导航路径
}

type Status struct {
	isRunning 		bool
	time 			int //开始模拟之后经过的时间（模拟时间）
	navigationList []Navi
}



var logicalToPoint map[string]Point //逻辑位置与地点的映射
var pointIndex map[Point]int //地点与图中索引值的映射
var graph []Vertex
