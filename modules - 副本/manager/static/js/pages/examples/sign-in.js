$(function () {
    $('#sign_in').validate({
        highlight: function (input) {
            console.log(input);
            $(input).parents('.form-line').addClass('error');
        },
        unhighlight: function (input) {
            $(input).parents('.form-line').removeClass('error');
        },
        errorPlacement: function (error, element) {
            $(element).parents('.input-group').append(error);
        }
    });

    $("#sign_in").submit(function () {
        $.ajax({
            url: "session",
            type: "GET",
            dataType:"json",
            data: {name: $("#sign_in [name=username]").val(),passwd: $("#sign_in [name=password]").val()},
            success: function (data) {
                try {
                    if (data.type == "sucess") {
                        location.href = "/default/index.html"
                        return
                    }
                    swal("", data.text, "info")
                }catch (e){
                    swal("", "服务接口错误", "error")
                }
            }, error: function () {
                swal("", "服务器链接错误", "error")
            }
        })
        return false
    })

});