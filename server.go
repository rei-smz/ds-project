package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("content-type","application/json")
	
	query := r.URL.Query()
	switch r.URL.Path {
		case "/api/check-service":
			w.Write([]byte("ok"))
		case "/api/add-student":
			stuName := query.Get("stuName")
			location := query.Get("location")
			Log("添加学生" + stuName + "在" + location)
			w.Write(AddStudent(stuName,location))
		case "/api/get-location":
			stuName := query.Get("stuName")
			jsonStr := "{\"location\":\""+locationInfo[stuName]+"\"}"
			w.Write([]byte(jsonStr))
		case "/api/get-facility-nearby":
			stuName := query.Get("stuName")
			res := GetNearby(stuName)
			data, _ := json.Marshal(res)
			w.Write(data)
		case "/api/add-navigation":
			var mustPass []string
			navStu := query.Get("navStu")
			dest := query.Get("dest")
			method, _ := strconv.Atoi(query.Get("method"))
			if method==2 {
				mustPass = strings.Split(query.Get("mustPass"),"-")
			} else {
				mustPass = []string{}
			}
			Log("为学生" + navStu + "添加导航，到" + dest)
			w.Write(AddNavi(navStu,dest,method,mustPass))
		case "/api/start-simulation":
			Log("模拟开始")
			wg.Add(1)
			go StartSimulation()
		case "/api/pause-simulation":
			if status.IsRunning {
				Log("模拟暂停")
				status.IsRunning = false
			}
		case "/api/get-navi-status":
			data, _ := json.Marshal(status)
			w.Write(data)
		case "/api/del-navi":
			navStu := query.Get("navStu")
			DelNavi(navStu)
	}
}

func Test() {
	for {
		res, err := http.Get("http://localhost:" + strconv.FormatInt(int64(port), 10) + "/api/check-service")
		if err == nil {
			defer res.Body.Close()
			body, _ := ioutil.ReadAll(res.Body)
			if string(body) == "ok" {
				Log("Web服务工作正常")
				break
			}
		}
	}
}

func Server() {
	http.HandleFunc("/api/", Handler)
	http.Handle("/", http.FileServer(http.Dir("static")))
	go Test()
	err := http.ListenAndServe(":" + strconv.FormatInt(int64(port), 10), nil)
	if err != nil {
		Log("Web服务未能正常启动")
		os.Exit(1)
	}
	wg.Done()
}
