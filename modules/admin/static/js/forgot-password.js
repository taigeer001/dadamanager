$(function () {
    $('#forgot_password').validate({
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

    $('#forgot_password').submit(function () {
        $.ajax({
            url: "forgot",
            type: "GET",
            dataType:"json",
            data: {mail: $("#forgot_password [name=email]").val(),name:$("#forgot_password [name=username]").val()},
            success: function (data) {
                try {
                    if (data.type == "success") {
                        location.href = "login.html"
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