/** wechat-pan-v1.0.0 LPPL License By http://www.coolcoder.cn */
;layui.define(["form", "upload"], function (e) {
    var i = layui.$, t = layui.layer, n = (layui.laytpl, layui.setter, layui.view, layui.admin), a = layui.form;
    layui.upload, i("body");
    a.render(), a.verify({
        nickname: function (e, i) {
            return new RegExp("^[a-zA-Z0-9_一-龥\\s·]+$").test(e) ? /(^\_)|(\__)|(\_+$)/.test(e) ? "用户名首尾不能出现下划线'_'" : /^\d+\d+\d$/.test(e) ? "用户名不能全为数字" : void 0 : "用户名不能有特殊字符"
        }, pass: [/^[\S]{6,12}$/, "密码必须6到12位，且不能出现空格"], repass: function (e) {
            if (e !== i("#LAY_password").val()) return "两次密码输入不一致"
        }
    }), e("set", {})
});