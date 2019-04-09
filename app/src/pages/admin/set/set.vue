<template>
  <div class="bell-set-wrapper">
    <bell-card>
      <div slot="card-header">
        修改密码
      </div>
      <div slot="card-content">
        <el-form ref="setForm" :model="form" :rules="rules" label-width="100px">
          <el-form-item label="当前密码" prop="old_password">
            <el-input v-model="form.old_password" type="password" />
          </el-form-item>
          <el-form-item label="新密码" prop="password">
            <el-input v-model="form.password" type="password" />
            <span class="form-word-aux">6到16个字符</span>
          </el-form-item>
          <el-form-item label="确认新密码" prop="confirmPsd">
            <el-input v-model="form.confirmPsd" type="password" @keyup.enter.native="onSubmit"/>
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
export default {
  components: {
    bellCard
  },
  data () {
    var validateOncePass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'));
      } else if (value !== this.form.password) {
        callback(new Error('两次输入密码不一致!'));
      } else {
        callback();
      }
    }
    return {
      form: {
        old_password: '',
        password: '',
        confirmPsd: ''
      },
      rules: {
        old_password: [
          { required: true, message: '请输入当前密码', trigger: 'blur' },
          { min: 6, max: 16, message: '长度在 6 到 16 个字符', trigger: 'blur' }
        ],
        password: [{ required: true, message: '请输入新密码', trigger: 'blur' }],
        confirmPsd: [{ required: true, validator: validateOncePass, trigger: 'blur' }]
      }
    }
  },
  methods: {
    onSubmit () {
      this.$refs['setForm'].validate((valid) => {
        if (valid) {
          loginAjax.resetPsd(this.form)
            .then(res => {
              this.resetForm()
              this.showMsg('success', '修改成功 ^_^')
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
    resetForm () {
      this.$refs['setForm'].resetFields();
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
.bell-set-wrapper
  .el-input
    width 30%
  .el-button--primary
    background-color: #009688;
    border-color #009688
    &:hover
      opacity .8
</style>
<style lang='stylus' scoped>
.bell-set-wrapper
  height: 100%;
  .form-word-aux
    color #999
    margin-left 5px
</style>
