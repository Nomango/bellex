<template>
  <div class="login-wrapper">
    <div class="bell-login-main">
      <div class="bell-login-header">
        <h2>智铃校园后台管理</h2>
        <el-form ref="form" :model="form">
          <el-form-item label="">
            <el-input
              v-model="form.username"
              type="text"
              placeholder="用户名">
              <i slot="prefix" class="iconfont icon-username"></i>
            </el-input>
          </el-form-item>
          <el-form-item label="">
            <el-input
              v-model="form.psd"
              type="password"
              placeholder="密码"
              @keyup.enter.native="onSubmit">
              <i slot="prefix" class="iconfont icon-password"></i>
            </el-input>
          </el-form-item>
          <el-form-item class="forget-item">
            <el-checkbox v-model="checked">记住密码</el-checkbox>
            <span class="forget-span">忘记密码?</span>
          </el-form-item>
          <el-form-item>
            <el-button type="primary"
              :loading="buttonLoading"
              @click="onSubmit">登入</el-button>
            <a :href="downloadUrl" target="_blank"><el-button style="margin-top:10px;">下载 PC 客户端</el-button></a>
          </el-form-item>
        </el-form>
      </div>
    </div>
    <div class="bell-login-footer">
      <p>© 2019 <a href="http://www.bellex.cn" target="_blank">bellex.cn</a></p>
      <p><a href="http://www.beian.miit.gov.cn" target="_blank">蜀ICP备16028738号-7</a></p>
    </div>
  </div>
</template>
<script>
import loginAjax from '@/api/login.js'
export default {
  data () {
    return {
      checked: false,
      buttonLoading: false,
      form: {
        username: '',
        psd: ''
      },
      downloadUrl: 'http://www.bellex.cn/download/latest'
    }
  },
  methods: {
    showMsg (msg, type) {
      this.$message({
        message: msg,
        type: type
      })
    },
    onSubmit () {
      this.buttonLoading = true
      if (this.form.username && this.form.psd) {
        loginAjax.getLogin({
          username: this.form.username,
          password: this.form.psd
        })
          .then(res => {
            if (res.data.redirect_url !== undefined) {
              this.showMsg('登入成功', 'success')
              window.location.href = res.data.redirect_url
            }
          })
          .catch(() => {
            this.buttonLoading = false
          })
      } else {
        this.showMsg('请输入必要字段', 'warning')
        setTimeout(() => {
          this.buttonLoading = false
        }, 500)
      }
    }
  }
}
</script>
<style lang="stylus">
.login-wrapper
  .el-button
    width 100%
    box-sizing border-box
  .el-button--primary
    background-color: #009688;
    border-color: #009688;
    &:hover
      opacity .8
  .el-input--prefix .el-input__inner
    padding-left 35px
  .el-input__prefix
    left 9px
  .el-form-item
    margin-bottom 14px
  .forget-item
    .el-form-item__content
      line-height 20px
    .el-checkbox__input.is-checked .el-checkbox__inner
      background-color: #5FB878;
      border-color: #5FB878;
    .el-checkbox__input.is-checked+.el-checkbox__label
      color #666
</style>
<style lang='stylus' scoped>
.login-wrapper
  position relative
  width 100%
  height 100%
  padding: 110px 0;
  min-height: 100%;
  box-sizing: border-box;
  .bell-login-main
    width: 375px;
    margin: 0 auto;
    box-sizing: border-box;
    .bell-login-header
      text-align center
      h2
        margin-bottom: 60px;
        font-weight: 500;
        font-size: 30px;
        color: #000;
    .contact-item
      line-height: 38px;
      text-align left
      span
        display: inline-block;
        vertical-align: middle;
        margin-right: 10px;
        font-size: 14px;
      .iconfont
        font-size 28px
        display: inline-block;
        vertical-align: middle;
        margin-right: 10px;
      .icon-qq
        color: #3492ED;
      .icon-weixin
        color: #4DAF29;
      .icon-weibo
        color: #CF1900;
    .forget-item
      text-align left
      text-indent 1px
      overflow hidden
      .forget-span
        float right
        cursor pointer
        color #029789
        &:hover
          opacity .8
  .bell-login-footer
    position: absolute;
    left: 0;
    bottom: 0;
    width: 100%;
    line-height: 30px;
    padding: 20px;
    text-align: center;
    box-sizing: border-box;
    color: rgba(0,0,0,.5);
    font-size 14px
    a
      padding: 0 5px;
      &:hover
        color rgba(0,0,0,1)
</style>
