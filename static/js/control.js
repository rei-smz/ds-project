var stuList = [];
var positionList = []; //应有一个读取位置的函数
var logicalList = [];
var mustPassPoint = []; //必经点
var logical
var navRunning,navStudent,navMode;
var port; //后端服务端口
var naviList;

navStudent = "选择学生…";
navMode = 0;

//这里写初始化函数
function Init(){
    //读取配置文件
    $.getJSON("../../config.json",function(jsonData) {
        port = jsonData["Port"];
        var points = jsonData["Point"];
        logical = jsonData["Logical"];
        $.each(points,function(i, value) {
            if(value.PointType == 0) {
                positionList.push(value.PointName);
            }
        });
        $.each(logical,function(i,value) {
            logicalList.push(value.LogicalName);
        });
    });
}

function SetNavStu(){
    navStudent = $("#choose_stu").children("option:selected").val();
};

function SetNavMethod(){
    var choose = $("#nav_method").children("option:selected").val();
    switch(choose){
        case "最短距离策略":
            navMode = 0;
            break;
        case "最短时间策略":
            navMode = 1;
            break;
        case "途径最短距离策略":
            navMode = 2;
            break;
        case "交通工具的最短时间策略":
            navMode = 3;
            break;
        default:
            break;
    }
};

function GetData(){ //模拟时读取后端数据函数
    var url = "http://localhost:" + String(port) + "/api/get-navi-status";
    $.getJSON(url,function(jsonData){
        naviList = new Array();
        naviList = jsonData["NavigationList"]
    });
}

function AddNav(){
    var dest = $("#nav_dest").val();
    if(navStudent == "选择学生…"){
        $("#info_content").text("请选择学生");
        $("#pop_info").modal("toggle");
        return;
    }
    if(positionList.indexOf(dest) == -1 && logicalList.indexOf(dest) == -1){
        $("#info_content").text("未设定目的地或目的地不在地图范围内！");
        $("#pop_info").modal("toggle");
        return;
    }
    if(logicalList.indexOf(dest) != -1 && dest != "食堂") {
        $("#output").append("<p>" + dest + "是一个逻辑位置，可导航到以下位置：</p>");
        $.each(logical,function(i,value) {
            if(value.LogicalName == dest) {
                $.each(value.PointName,function(i,pname) {
                    $("#output").append("<p>" + pname + "</p>");
                });
            }
        });
        return;
    }
    if(navMode == 2){
        var must_pass = $("#must_pass").val()
        mustPassPoint = must_pass.split(" ");
        if(mustPassPoint==[]){
            $("#info_content").text("必经位置不能为空！");
            $("#pop_info").modal("toggle");
            return;
        }
        for(var i in mustPassPoint){
            if(!(i in positionList)){
                $("#info_content").text("必经位置不在可导航的位置内！");
                $("#pop_info").modal("toggle");
                return;
            }
        }
        var mustPassStr = mustPassPoint.join("-");
        //向后端发送数据
        $.get("http://localhost:" + String(port) + "/api/add-navigation",{navStu: navStudent,dest: dest, method: navMode,mustPass: mustPassStr}, function(data) {
            console.log(data);
        });
         $("#info_content").text("添加导航成功！");
        $("#pop_info").modal("toggle");
        if($("#progress" + String(stuList.indexOf(navStudent))).length != 0) {
            return;
        }
        $("#progressbar_area").append("<div class=\"progress\" style=\"margin-top: 10px;\" id=\"progress"+String(stuList.indexOf(navStudent))+"\"><div class=\"progress-bar progress-bar-info\" role=\"progressbar\" aria-valuemin=\"0\" aria-valuemax=\"100\" style=\"width: 0%;\">学生"+navStudent+"行程进度</div></div>");
    }
    else{
        //向后端发送数据
        $.get("http://localhost:" + String(port) + "/api/add-navigation",{navStu: navStudent,dest: dest, method: navMode}, function(data) {
            console.log(data);
        });
        $("#info_content").text("添加导航成功！");
        $("#pop_info").modal("toggle");
        if($("#progress" + String(stuList.indexOf(navStudent))).length != 0) {
            return;
        }
        $("#progressbar_area").append("<div class=\"progress\" style=\"margin-top: 10px;\" id=\"progress"+String(stuList.indexOf(navStudent))+"\"><div class=\"progress-bar progress-bar-info\" role=\"progressbar\" aria-valuemin=\"0\" aria-valuemax=\"100\" style=\"width: 0%;\">学生"+navStudent+"行程进度</div></div>");
    }
    //此处创建进度条
}

function StartNav(){ //开始导航控件
    var callback;
    callback = $.get("http://localhost:" + String(port) + "/api/start-simulation",function(data){
        return data;
    });
    $("#nav_control").removeClass("btn-success");
    $("#nav_control").addClass("btn-danger");
    $("#nav_control").text("暂停模拟 ❚❚");
    $("#nav_control").attr("onclick", "PauseNav()");
    $("#add_nav").hide();
    $("#get_facility_nearby").hide();
    $("#choose_stu").attr("disabled",true);
    $("#nav_dest").attr("disabled",true);
    $("#nav_method").attr("disabled",true);
    $("#must_pass").attr("disabled", true);
    $("#get_location").hide();
    $("#output").append("<p>开始模拟</p>");
    GetData();
    navRunning = setInterval(function(){ //开始模拟后每秒读取一次数据
        GetData();
        //这里应有对数据的处理
        //e.g 通过获取到的数据更新进度条
        var $res = $("#progressbar_area").find(".progress");
        var delList = []
        //console.log($res);
        for(var i=0;i<$res.length;i++) {
            //console.log("naviList len="+naviList.length);
            for(var ni=0; ni < naviList.length; ni ++) {
                if("progress"+stuList.indexOf(naviList[ni]["StudentName"])==$res.eq(i).attr("id")) {
                    console.log(naviList[ni]["Distance"]);
                    if(naviList[ni]["Distance"] <= 0) {
                        $("#output").append("<p>学生" + naviList[ni]["StudentName"] + "行程结束</p>");
                        $.get("http://localhost:" + String(port) + "/api/del-navi", {navStu: naviList[ni]["StudentName"]});
                        GetData();
                        delList.push(i);
                    } else {
                        var $bar = $res.eq(i).children("div").eq(0)
                        $bar.attr("style","width:"+String(naviList[ni]["Time"]/naviList[ni]["RemainingTime"]*100)+"%");
                    }
                }
            }
        }
        for(var i =0; i < delList.length; i ++){
            $res.eq(delList[i]).remove();
        }
    },1000);
}

function PauseNav(){ //暂停导航控件
    var callback = $.get("http://localhost:" + String(port) + "/api/pause-simulation",function(data){
        return data;
    });
    clearInterval(navRunning);
    $("#nav_control").removeClass("btn-danger");
    $("#nav_control").addClass("btn-success");
    $("#nav_control").text("开始模拟 ➤");
    $("#nav_control").attr("onclick", "StartNav()");
    $("#add_nav").show();
    $("#get_facility_nearby").show();
    $("#choose_stu").removeAttr("disabled");
    $("#nav_dest").removeAttr("disabled");
    $("#nav_method").removeAttr("disabled");
    $("#must_pass").removeAttr("disabled");
    $("#get_location").show();
    $("#output").append("<p>模拟暂停</p>");
}

function GetFacility(){ //获取周边设施
    if(navStudent == "选择学生…"){
        $("#info_content").text("请选择学生");
        $("#pop_info").modal("toggle");
        return;
    }
    $.getJSON("http://localhost:" + String(port) + "/api/get-facility-nearby", {stuName: navStudent},function(data){
        var nearbyList = data["nearby"];
        if(nearbyList == null) {
            $("#output").append("<p>学生" + navStudent + "周边无可用的设施</p>");
        } else {
            $("#output").append("<p>学生" + navStudent + "周边的设施有：</p>");
            $.each(nearbyList, function(i, value) {
                $("#output").append("<p>" + value.NearbyName + "，距离当前位置" + String(value.Dis) + "米</p>");
            });
        }
    });
}

function GetLocation() {
    if(navStudent == "选择学生…"){
        $("#info_content").text("请选择学生");
        $("#pop_info").modal("toggle");
        return;
    }
    $.getJSON("http://localhost:" + String(port) + "/api/get-location", {stuName: navStudent},function(data) {
        var location = data["location"];
        if(location == null) {
            $("#output").append("<p>未找到学生"+navStudent+"的当前位置</p>");
        } else {
            $("#output").append("<p>学生"+navStudent+"当前在"+location+"</p>");
        }
    });
}

function ClearOutput(){
    $("#output").empty();
}

function AddStudent(){
    var studentName = $("#stu_name").val();
    var position = $("#stu_position").val();
    if(studentName=="")
    {
        $("#info_content").text("学生姓名不能为空！");
        $("#pop_info").modal("toggle");
        return;
    }
    if(stuList.indexOf(studentName)!=-1)
    {
        $("#info_content").text("不能重复添加学生！");
        $("#pop_info").modal("toggle");
        return;
    }
    if(positionList.indexOf(position)==-1)
    {
        $("#info_content").text("未设定地点或设定的地点不在地图范围内！");
        $("#pop_info").modal("toggle");
        return;
    }
    //这里向后端发送数据
    $.get("http://localhost:"+String(port)+"/api/add-student",{stuName: studentName,location: position},function(data) {
        console.log(data); 
    });
    $("#info_content").text("添加学生成功！");
    $("#pop_info").modal("toggle");
    stuList.push(studentName);
    $("#choose_stu").append("<option>"+studentName+"</option>");
}

$(function() {
    Init();
});
