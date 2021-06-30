package main

import (
	"errors"
)

var locationInfo map[string]string //学生位置信息，格式为["名字"]: "位置"

type Point struct {
	PointType 	int //地点类型，0为设施、地标或楼层，1为路口
	schoolId 	int //校区编号
	PointName 		string
}

type Road struct {
	StartName 	   string
	EndName        string
	Length         float64
	IsBikeRoad     bool
	//crowdedForWalk float32 //步行的拥挤度
	//crowdedForBike float32 //骑车的拥挤度
}

type Logical struct {
	LogicalName string
	PointName 	[]string
}

type Arc struct {
	targetIndex     int
	length          float64
	isBikeRoad      bool
	crowdedForWalk  float32 //步行的拥挤度
	crowdedForBike  float32 //骑车的拥挤度
	timeWalk 	    float64
	timeBike 	    float64
}

type Vertex struct {
	name 		string
	pointType 	int
	number 		int //与该点有关的边的个数
	schoolId 	int
	book 		bool //每次查询都需初始化,标记更新dis时是否被访问
	pre 		int //每次查询都需初始化,记录其前驱
	dis 		float64 //每次查询都需初始化,表示与源点的最短距离/最小时间
	isNeeded 	bool //是否为必经点
	preArc 		int //记录边的编号，即使用了前一个点的哪条边来到达该点，可与pre同时记录
	arcList []Arc
}

type Trip struct {
	StartIndex 	int //当前行程路段的起点（使用图的顶点编号）
	ArcIndex 	int //当前行程路段走的边（使用起点边列表中的编号）
	Length 		float64 //路段长度，规划路径时赋值
	Passed 		float64 //路段上的已走距离
	StartTime 	float64
	Time 		float64 //需要时间
}

type Navi struct {
	StudentName 	string
	Distance 		float64 //目的地距离
	DestName 		string //目的地名称
	StartTime 		float64 //开始时间，添加导航时设为status.time
	Time 			float64 //已用时间，每次更新导航状态时赋值为status.time-startTime
	RemainingTime 	float64 //剩余时间
	CurrentIndex 	int //当前走到的path中的路段编号
	Path 			[]Trip //导航路径
}

type Status struct {
	IsRunning 		bool
	Time 			float64 //开始模拟之后经过的时间（模拟时间）
	NavigationList []Navi
}

type nearbyInfo struct {
	NearbyName string
	Dis 		float64
}

var logicalToPoint map[string][]string //逻辑位置与地点的映射
var pointIndex map[Point]int //地点与图中索引值的映射
var pointNameIndex map[string]int //点名称与图中索引值的映射

/***********栈的实现***********/

type Stack []interface{}

func (stack Stack) Len() int {
	return len(stack)
}
func (stack Stack) Cap() int {
	return cap(stack)
}
func (stack *Stack) Push(value interface{})  {
	*stack = append(*stack, value)
}
func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("empty")
	}
	return stack[len(stack) - 1], nil
}
func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return nil, errors.New("empty")
	}
	value := theStack[len(theStack) - 1]
	*stack = theStack[:len(theStack) - 1]
	return value, nil
}
func (stack Stack) isEmpty() bool {
	return len(stack) == 0
}

var graph []Vertex
