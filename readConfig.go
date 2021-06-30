package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ConfigFile struct {
	Port int
	Point []Point
	Logical []Logical
}

type RoadFile struct {
	Road []Road
}

func ReadConfig()  {
	raw, err := ioutil.ReadFile("static/config.json")
	if err != nil {
		Log("打开配置文件失败")
		os.Exit(1)
	}
	raw1, err := ioutil.ReadFile("static/roads.json")
	if err != nil {
		Log("打开配置文件失败")
		os.Exit(1)
	}
	var config ConfigFile
	var roadConfig RoadFile
	err = json.Unmarshal(raw, &config)
	if err != nil {
		Log("解析配置文件失败")
		os.Exit(1)
	}
	err = json.Unmarshal(raw1, &roadConfig)
	if err != nil {
		Log("解析配置文件失败")
		os.Exit(1)
	}
	port = config.Port
	pointList = config.Point
	roadList = roadConfig.Road
	logicalList = config.Logical
	for _, i := range logicalList {
		logicalToPoint[i.LogicalName] = i.PointName
	}
	Log(fmt.Sprintf("成功读取%d个点，%d条边，%d条逻辑位置信息", len(pointList), len(roadList), len(logicalList)))
}
