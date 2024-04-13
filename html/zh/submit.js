function submit() {
    var mcid = $("#mcid").val();
    var contact_id = $("#contact_id").val();
    if (mcid == "") {
        alertify.alert("提醒","请输入Minecraft ID");
        return;
    }
    if (contact_id == "") {
        alertify.alert("提醒","请输入QQ号");
        return;
    }


    $.ajax({
        type: "GET",
        url: "/api/applywhitelist",
        data: {
            name: mcid,
            contactid: contact_id
        },
        dataType: 'json',
        success: function(data) {
            console.log(data);
            if (data["success"]){
                alertify.alert("提醒","您的申请已经提交，请耐心等待审核哦！");
            } else {
                alertify.alert("提醒","未知错误！请稍后再试");
            }
        },
        error: function(data) {
            console.log(data);
            console.log(data.responseText);
            var j = JSON.parse(data.responseText);
            alertify.alert("提醒", j["error"]);
        }
    });
}