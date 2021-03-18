var stuList = []

function StartNav(){ //开始导航控件
    //这里应向后端通信
    $("#nav_control").removeClass("btn-success");
    $("#nav_control").addClass("btn-danger");
    $("#nav_control").text("暂停导航 ❚❚");
    $("#nav_control").attr("onclick", "PauseNav()");
    $("#get_facility_nearby").hide();
}

function PauseNav(){ //暂停导航控件
    //这里应向后端通信
    $("#nav_control").removeClass("btn-danger");
    $("#nav_control").addClass("btn-success");
    $("#nav_control").text("开始导航 ➤");
    $("#nav_control").attr("onclick", "StartNav()");
    $("#get_facility_nearby").show();
}

function GetFacility(){ //获取周边设施
    $("#output").append("<p>获取设施测试</p>");
}

function ClearOutput(){
    $("#output").empty();
}
