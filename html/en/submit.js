function submit() {
    var mcid = $("#mcid").val();
    var contact_method = $("#contact_method").val();
    var contact_id = $("#contact_id").val();
    if (mcid == "") {
        alertify.alert("Caution","Please type in your Minecraft ID");
        return;
    }
    if (contact_method == "") {
        alertify.alert("Caution","Please type in your contact method");
        return;
    }
    if (contact_id == "") {
        alertify.alert("Caution","Please type in your contact ID");
        return;
    }


    $.ajax({
        type: "GET",
        url: "/api/applywhitelist",
        data: {
            name: mcid,
            contactmethod: contact_method,
            contactid: contact_id
        },
        dataType: 'json',
        success: function(data) {
            console.log(data);
            if (data["success"]){
                alertify.alert("Caution","Your application has been submitted successfully. Please wait for the admin to review it.");
            } else {
                alertify.alert("Caution","Unknown error, please contact admin.");
            }
        },
        error: function(data) {
            console.log(data);
            console.log(data.responseText);
            var j = JSON.parse(data.responseText);
            alertify.alert("Caution", j["error"]);
        }
    });
}