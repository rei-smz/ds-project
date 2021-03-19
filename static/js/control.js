var stuList = [];
var positionList = ["aaa"]; //应有一个读取位置的函数
var navRunning,navStudent,navMode;

navStudent = "选择学生…";
navMode = 0;

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

}

function StartNav(){ //开始导航控件
    var dest = $("#nav_dest").val();
    if(navStudent == "选择学生…"){
        $("#info_content").text("请选择学生");
        $("#pop_info").modal("toggle");
        return;
    }
    if(positionList.indexOf(dest) == -1){
        $("#info_content").text("未设定目的地或目的地不在地图范围内！");
        $("#pop_info").modal("toggle");
        return;
    }
    //这里应向后端通信
    $("#nav_control").removeClass("btn-success");
    $("#nav_control").addClass("btn-danger");
    $("#nav_control").text("暂停模拟 ❚❚");
    $("#nav_control").attr("onclick", "PauseNav()");
    $("#get_facility_nearby").hide();
    $("#choose_stu").attr("disabled",true);
    $("#nav_dest").attr("disabled",true);
    $("#nav_method").attr("disabled",true);
    navRunning = setInterval(function(){ //开始模拟后每秒读取一次数据
        GetData();
        //这里应有对数据的处理
        $("#output").append("<p>"+navStudent+" "+navMode+"</p>");
    },1000);
}

function PauseNav(){ //暂停导航控件
    //这里应向后端通信
    clearInterval(navRunning);
    $("#nav_control").removeClass("btn-danger");
    $("#nav_control").addClass("btn-success");
    $("#nav_control").text("开始模拟 ➤");
    $("#nav_control").attr("onclick", "StartNav()");
    $("#get_facility_nearby").show();
    $("#choose_stu").removeAttr("disabled");
    $("#nav_dest").removeAttr("disabled");
    $("#nav_method").removeAttr("disabled");
}

function GetFacility(){ //获取周边设施
    $("#output").append("<p>获取设施测试</p>");
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
        return;
    }
    if(stuList.indexOf(studentName)!=-1)
    {
        $("#info_content").text("不能重复添加学生！");
        return;
    }
    if(positionList.indexOf(position)==-1)
    {
        $("#info_content").text("未设定地点或设定的地点不在地图范围内！");
        return;
    }
    $("#info_content").text("成功添加学生！");
    stuList.push(studentName);
    $("#choose_stu").append("<option>"+studentName+"</option>");
    //这里向后端发送数据
}
