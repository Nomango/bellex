<template>
  <div class="tableDetail-wrapper">
    <bell-card>
      <div class="card-header-content" slot="card-header">
        <el-button type="success" @click="handleBack">返回</el-button>
      </div>
      <div class="card-content-body" slot="card-content">
        <el-form ref="detailForm" :model="detailForm" label-width="80px">
          <el-form-item label="名称">
            <el-input v-model="detailForm.name"></el-input>
          </el-form-item>
          <el-form-item label="识别码">
            <el-input v-model="detailForm.code"></el-input>
          </el-form-item>
          <el-form-item label="密码">
            <el-input v-model="detailForm.psd"></el-input>
          </el-form-item>
          <el-form-item label="时间表">
            <el-select v-model="detailForm.time" placeholder="请选择时间">
              <el-option label="2018-08" value="2018-08"></el-option>
              <el-option label="2018-09" value="2018-09"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item class="footer-item">
            <el-button type="primary" @click="onSubmit">确定</el-button>
          </el-form-item>
        </el-form>
      </div>
    </bell-card>
  </div>
</template>
<script>
import bellCard from 'common/card/card'
import homeAjax from '@/api/home.js'
export default {
  components: {
    bellCard
  },
  data () {
    return {
      detailForm: {
        name: '',
        code: '',
        psd: '',
        time: ''
      }
    }
  },
  created () {
    console.log('ss', this.$route)
    this.getdetailData()
  },
  methods: {
    getdetailData () {
      homeAjax.getControllers({})
        .then(res => {
          console.log('res', res)
          if (res.code === 0) {
            res = res.data
            let filterData = []
            filterData = res.controllerList.filter(item => item.id === this.$route.params.id)
            console.log(filterData, 'aaa')
            this.detailForm = Object.assign({}, filterData[0])
          }
        })
        .catch(err => {
          console.log('err', err)
        })
    },
    handleBack () {
      this.$router.go(-1)
    },
    onSubmit () {
      console.log('sss')
    }
  }
}
</script>
<style lang="stylus">
.tableDetail-wrapper
  .el-button--success
    background-color: #009688;
    border-color #009688
    &:hover
      opacity: .8;
      color: #fff;
  .el-select
    width 100%
  .footer-item
    .el-button--primary
      height: 38px;
      line-height: 38px;
      padding: 0 18px;
      background-color: #009688;
      border-color #009688
      &:hover
        opacity .8
</style>
<style lang='stylus' scoped>
.tableDetail-wrapper
  position: relative;
  height 100%
  margin: 0 auto;
  box-sizing border-box
  .page-wrapper
    margin-top 20px
  .card-content-body
    width 400px
    margin 0 auto
</style>
