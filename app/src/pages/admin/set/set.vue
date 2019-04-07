<template>
  <div class="bell-set-wrapper">
    <bell-card>
      <div slot="card-header">
        修改密码
      </div>
      <div slot="card-content">
        <el-form ref="setForm" :model="form" :rules="rules" label-width="100px">
          <el-form-item label="当前密码" prop="psd">
            <el-input v-model="form.psd" />
          </el-form-item>
          <el-form-item label="新密码" prop="newPsd">
            <el-input v-model="form.newPsd" />
            <span class="form-word-aux">6到16个字符</span>
          </el-form-item>
          <el-form-item label="确认新密码" prop="confirmPsd">
            <el-input v-model="form.confirmPsd" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="onSubmit">立即创建</el-button>
          </el-form-item>
        </el-form>
      </div>
    </bell-card>
  </div>
</template>
<script>
import bellCard from 'common/card/card'
export default {
  components: {
    bellCard
  },
  data () {
    return {
      form: {
        psd: '',
        newPsd: '',
        confirmPsd: ''
      },
      rules: {
        psd: [
          { required: true, message: '请输入活动名称', trigger: 'blur' },
          { min: 6, max: 16, message: '长度在 6 到 16 个字符', trigger: 'blur' }
        ],
        newPsd: [{ required: true, message: '请输入新密码', trigger: 'blur' }],
        confirmPsd: [{ required: true, message: '请确认密码', trigger: 'blur' }]
      }
    }
  },
  methods: {
    onSubmit () {
      this.$refs['setForm'].validate((valid) => {
        if (valid) {
          console.log('submit!')
        } else {
          console.log('error submit!!')
          return false
        }
      });
      console.log('aaa')
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
