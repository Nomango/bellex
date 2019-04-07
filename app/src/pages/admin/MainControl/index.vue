<template>
  <div class="mainControl-wrapper">
    <bell-card>
      <div class="card-header-content" slot="card-header">
        <el-button type="success">添加</el-button>
      </div>
      <div class="card-content-body" slot="card-content">
          <el-table
            :data="controllerData"
            stripe
            border
            style="width: 100%">
          <el-table-column
            v-for="(item,index) of columns"
            :key="index"
            :prop="item.prop"
            align="center"
            :label="item.label" />
          <el-table-column
            prop="address"
            label="操作"
            align="center"
            width="200">
            <template slot-scope="scope">
              <el-button @click="handleCheck(scope.row)" icon="el-icon-edit" size="mini">查看</el-button>
              <el-button @click="handleDelete(scope.row)" icon="el-icon-delete" type="danger" size="mini">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
        <div class="page-wrapper">
          <bell-pagination
            :total="totalPage"
            :currentPage="currentPage"
            :pageSizes="pageSizes"
            @sizeChange="handleSizeChange"
            @currentChange="handleCurrentChange" />
        </div>
        <div class="normal-dialog">
          <bell-dialog
            width="24%"
            :dialogVisible="controlDialog"
            @confirm="handleConfirm"
            @cancel="handleCancel" />
        </div>
      </div>
    </bell-card>
  </div>
</template>
<script>
import bellCard from 'common/card/card'
import bellPagination from 'common/pagination/pagination'
import bellDialog from 'common/dialog/dialog'
import homeAjax from '@/api/home.js'
export default {
  components: {
    bellCard,
    bellPagination,
    bellDialog
  },
  data () {
    return {
      currentPage: 1,
      totalPage: 0,
      pageSizes: {
        sizeArr: [5, 8, 10, 20],
        size: 5
      },
      controlDialog: false,
      controllerData: [],
      columns: [{
        prop: 'name',
        label: '名称'
      }, {
        prop: 'network_state',
        label: '网络状态'
      }, {
        prop: 'idle_state',
        label: '待机状态'
      }, {
        prop: 'create_time',
        label: '创建时间'
      }]
    }
  },
  mounted () {
    this.initController()
  },
  methods: {
    initController () {
      homeAjax.getControllers({})
        .then(res => {
          console.log('res', res)
          if (res.code === 0) {
            res = res.data
            this.controllerData = res.controllerList
            this.totalPage = res.total
          } else {
            this.controllerData = []
          }
        })
        .catch(err => {
          console.log('err', err)
        })
    },
    handleCheck (val) {
      console.log('val', val)
      this.$router.push({ name: 'tableDetail', params: val.id })
    },
    handleDelete (val) {
      this.controlDialog = true
      console.log('handleDelete', val)
    },
    handleSizeChange (val) {
      console.log('handleSizeChange', val)
    },
    handleCurrentChange (val) {
      console.log('handleCurrentChange', val)
    },
    handleConfirm (val) {
      this.controlDialog = val
    },
    handleCancel (val) {
      this.controlDialog = val
    }
  }
}
</script>
<style lang="stylus">
.mainControl-wrapper
  .el-button--success
    background-color: #009688;
    border-color #009688
    &:hover
      opacity: .8;
      color: #fff;
  .normal-dialog
    .el-dialog__header
      padding: 0 80px 0 20px;
      height: 42px;
      line-height: 42px;
      border-bottom: 1px solid #eee;
      font-size: 14px;
      color: #333;
      overflow: hidden;
      background-color: #F8F8F8;
      border-radius: 2px 2px 0 0;
    .el-dialog__headerbtn
      top: 13px;
      right: 17px;
    .el-dialog__body
      padding 20px
    .el-dialog__footer
      .el-button
        height: 28px;
        line-height: 28px;
        margin: 5px 5px 0;
        padding: 0 15px;
        border-radius: 2px;
        font-weight: 400;
</style>
<style lang='stylus' scoped>
.mainControl-wrapper
  position: relative;
  height 100%
  margin: 0 auto;
  box-sizing border-box
  .page-wrapper
    margin-top 20px
</style>
