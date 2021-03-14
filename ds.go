package main

/*type Location struct {
	name string
	locationName string
}*/

var locationInfo map[string]string //学生位置信息，格式为["名字"]: "位置"

type Point struct {
	pointType int //地点类型
	name string
	pointsNearby []Point //附近的建筑
	roadList []Road
}

type Road struct {
	targetName     string
	length         int
	isBikeRoad     bool
	crowdedForWalk float64 //步行的拥挤度
	crowdedForBike float64 //骑车的拥挤度
}

type Navi struct {
	studentName string
	distance int //目的地距离
	remainingTime int
}

type Status struct {
	isRunning bool
	time int
	navigationList []Navi
}
