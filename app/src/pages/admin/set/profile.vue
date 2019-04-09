<template>
  <div class="bell-profile-wrapper">
    <bell-card>
      <div slot="card-header">
        修改资料
      </div>
      <div slot="card-content">
        <el-form ref="setForm" :model="form" :rules="rules" label-width="100px">
          <el-form-item label="用户名" prop="username">
            <el-input v-model="form.username" />
          </el-form-item>
          <el-form-item label="昵称" prop="nickname">
            <el-input v-model="form.nickname" />
          </el-form-item>
          <el-form-item label="邮箱" prop="email">
            <el-input v-model="form.email" @keyup.enter.native="onSubmit"/>
          </el-form-item>
          <el-form-item>
            <el-button
              type="primary"
              @click="onSubmit">确定</el-button>
          </el-form-item>
        </el-form>
      </div>
    </bell-card>
  </div>
</template>
<script>
import bellCard from 'common/card/card'
import loginAjax from '@/api/login.js'
import { mapMutations } from 'vuex'
export default {
  components: {
    bellCard
  },
  data () {
    return {
      form: {
        username: '',
        nickname: '',
        email: ''
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' }
        ],
        nickname: [{ required: true, message: '请输入昵称', trigger: 'blur' }],
        email: [{ required: true, message: '请输入邮箱', trigger: 'blur' }]
      }
    }
  },
  created () {
    this.getProfileInfo()
  },
  methods: {
    ...mapMutations(['SET_NICKNAME']),
    getProfileInfo () {
      loginAjax.getProfileInfo()
        .then(res => {
          res = res.data
          this.form = Object.assign({}, res.user)
        })
        .catch(err => {
          console.log(err)
        })
    },
    onSubmit () {
      this.$refs['setForm'].validate((valid) => {
        if (valid) {
          loginAjax.updateProfileInfo(this.form)
            .then(res => {
              this.showMsg('success', '修改成功 ^_^')
              this.SET_NICKNAME(this.form.nickname)
            })
            .catch(err => {
              console.log('err', err)
            })
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    showMsg (type, msg) {
      this.$message({
        type: type,
        message: msg
      })
    }
  }
}
</script>
<style lang="stylus">
.bell-profile-wrapper
  .el-input
    width 30%
  .el-button--primary
    background-color: #009688;
    border-color #009688
    &:hover
      opacity .8
</style>
<style lang='stylus' scoped>
.bell-profile-wrapper
  height: 100%;
</style>
