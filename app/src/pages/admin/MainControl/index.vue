<template>
  <div class="mainControl-wrapper">
    <bell-card>
      <div class="card-header-content" slot="card-header">
        <el-button type="success" icon="el-icon-plus" @click="addClick">添加</el-button>
        <el-button type="success" icon='el-icon-refresh' @click="refreshClick">刷新</el-button>
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
            prop="schedule"
            align="center"
            label="时间表">
            <template slot-scope="{row}">
              <span>{{row.schedule.name}}</span>
            </template>
          </el-table-column>
          <el-table-column
            align="center"
            label="网络状态">
            <template slot-scope="{row}">
              <el-tag v-if="!row.accept" :type="row.accept | tagsFilter">{{row.accept | acceptFilter}}</el-tag>
              <template v-else>
                <el-tooltip  content="点我断开" placement="top">
                  <el-button
                    type="success"
                    size="mini"
                    @click="handleAccept(row)">{{row.accept | acceptFilter}}</el-button>
                </el-tooltip>
              </template>
            </template>
          </el-table-column>
          <el-table-column
            align="center"
            label="待机状态">
            <template slot-scope="{row}">
              <el-tag :type="row.idle | tagsFilter">{{row.idle | idleFilter}}</el-tag>
            </template>
          </el-table-column>
          <el-table-column
            label="操作"
            align="center"
            width="300">
            <template slot-scope="scope">
              <el-button @click="handleEdit(scope.row)" icon="el-icon-edit" size="mini">编辑</el-button>
              <el-dropdown @command="handleCommand(scope.row, $event)">
                <el-button icon="el-icon-phone-outline" type="success" size="mini">打铃</el-button>
                <el-dropdown-menu slot="dropdown">
                  <el-dropdown-item command="a">立即打铃</el-dropdown-item>
                  <el-dropdown-item command="b">定时打铃</el-dropdown-item>
                </el-dropdown-menu>
              </el-dropdown>
              <el-button @click="handleDelete(scope.row)" icon="el-icon-delete" type="danger" size="mini">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
        <div class="page-wrapper">
          <bell-pagination
            v-if="totalPage > 0"
            :total="totalPage"
            :currentPage="currentPage"
            :pageSizes="pageSizes"
            @sizeChange="handleSizeChange"
            @currentChange="handleCurrentChange" />
        </div>
        <div class="add-dialog">
          <add-dialog
            width="40%"
            :title="clickType[clickStatus]"
            :dialogVisible="addDialog"
            @cancel="handleAddCancel"
            @confirm="handleAddConfirm">
            <div class="content" slot="content">
              <el-form ref="controlForm" :model="controlForm" :rules="rules" label-width="70px">
                <el-form-item label="名称: " prop="name">
                  <el-input v-model="controlForm.name" placeholder="请输入主控机名称" />
                </el-form-item>
                <el-form-item label="序列号: " prop="code">
                  <el-input :disabled="clickStatus === 'update'" v-model="controlForm.code" placeholder="请输入主控机序列号" />
                </el-form-item>
                <el-form-item label="时间表: " prop="schedule_id">
                  <el-select v-model="controlForm.schedule_id" placeholder="请选择应用到主控机的时间表">
                    <el-option
                      v-for="(item,index) of timeOption"
                      :key="index"
                      :label="item.label"
                      :value="item.value" />
                  </el-select>
                </el-form-item>
              </el-form>
            </div>
          </add-dialog>
        </div>
        <el-dialog
          title="设置打铃时间"
          :modal-append-to-body="false"
          :visible.sync="timingFormVisible"
          width="30%">
          <el-form
            ref="timingForm"
            :rules="rules"
            :model="timingTemp"
            label-position="right"
            label-width="70px"
            style="width: 80%; margin-left: 5%;">
            <el-form-item label="时间: " prop="time">
              <el-time-picker
                v-model="timingTemp.time"
                format="HH:mm"
                :picker-options="{
                  selectableRange: '00:00:00 - 23:59:59'
                }"
                placeholder="设置打铃时间" />
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button type="primary" @click="confirmClick">确定</el-button>
          </div>
        </el-dialog>
      </div>
    </bell-card>
  </div>
</template>
<script>
import bellCard from 'common/card/card'
import bellPagination from 'common/pagination/pagination'
import addDialog from 'common/dialog/addDialog'
import { translateTime } from '@/utils/tools.js'
import homeAjax from '@/api/home.js'
export default {
  components: {
    bellCard,
    bellPagination,
    addDialog
  },
  filters: {
    idleFilter (status) {
      const statusMap = {
        true: '就绪',
        false: '工作中'
      }
      return statusMap[status]
    },
    acceptFilter (status) {
      const statusMap = {
        true: '已连接',
        false: '已断开'
      }
      return statusMap[status]
    },
    tagsFilter (status) {
      const statusMap = {
        false: 'info',
        true: 'success'
      }
      return statusMap[status]
    }
  },
  data () {
    return {
      addDialog: false,
      timingFormVisible: false,
      currentPage: 1,
      totalPage: 0,
      pageSizes: {
        sizeArr: [5, 8, 10, 20],
        size: 5
      },
      controllerData: [],
      columns: [{
        prop: 'name',
        label: '名称'
      }, {
        prop: 'code',
        label: '序列号'
      }, {
        prop: 'secret',
        label: '密码'
      }],
      timeOption: [],
      controlForm: {
        code: '',
        schedule_id: '',
        name: ''
      },
      clickType: {
        update: '编辑',
        create: '添加'
      },
      clickStatus: null,
      rules: {
        name: [{ required: true, message: '请选择名称', trigger: 'blur' }],
        code: [
          { required: true, message: '请输入主控机编号', trigger: 'blur' },
          { min: 8, max: 8, message: '长度为 8 个数字字母的组合字符', trigger: 'blur' }
        ],
        schedule_id: [{ required: true, message: '请选择时间', trigger: 'change' }],
        time: [{ required: true, message: '请选择时间', trigger: 'change' }]
      },
      timingTemp: {
        time: ''
      },
      timingId: null
    }
  },
  created () {
    this.iniData()
  },
  methods: {
    iniData () {
      this.initController()
      this.initTime()
    },
    initTime () {
      homeAjax.getTimeList({
        page: 0,
        limit: 0
      })
        .then(res => {
          console.log('time', res)
          res = res.data
          this.handleTimeData(res.data)
        })
        .catch(err => {
          console.log(err)
        })
    },
    handleTimeData (res) {
      this.timeOption = res.map(item => {
        return {
          label: item.name,
          value: item.id
        }
      })
    },
    initController () {
      homeAjax.getControllers({
        page: this.currentPage,
        limit: this.pageSizes.size
      })
        .then(res => {
          console.log('res', res)
          res = res.data
          this.controllerData = res.data
          this.totalPage = res.total
        })
        .catch(err => {
          console.log('err', err)
        })
    },
    handleEdit (val) {
      this.controlForm = Object.assign({}, val)
      this.controlForm.schedule_id = val.schedule.id
      this.clickStatus = 'update'
      this.addDialog = true
      this.resetForm()
    },
    handleDelete (val) {
      this.$confirm('此操作将永久删除, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        homeAjax.delControllers({
          id: val.id
        })
          .then(res => {
            this.showMsg('success', '删除成功!')
            this.iniData()
          })
          .catch(err => {
            console.log(err)
          })
      }).catch(() => {
        this.showMsg('info', '已取消删除')
      })
    },
    handleSizeChange (val) {
      console.log('handleSizeChange', val)
      this.pageSizes.size = val
      this.initController()
    },
    handleCurrentChange (val) {
      console.log('handleCurrentChange', val)
      this.currentPage = val
      this.initController()
    },
    addClick () {
      this.resetFormData()
      this.clickStatus = 'create'
      this.addDialog = true
      this.resetForm()
    },
    createData () {
      homeAjax.addControllers(this.controlForm)
        .then(res => {
          this.showMsg('success', '添加成功 ^_^')
          this.initController()
          this.addDialog = false
        })
        .catch(err => {
          console.log(err)
        })
    },
    updateData () {
      homeAjax.putControllers(this.controlForm)
        .then(res => {
          this.showMsg('success', '更新成功 ^_^')
          this.initController()
          this.addDialog = false
        })
        .catch(err => {
          console.log(err)
        })
    },
    handleAddCancel () {
      this.addDialog = false
    },
    handleAddConfirm () {
      this.$refs['controlForm'].validate((valid) => {
        if (valid) {
          this.clickStatus === 'create' ? this.createData() : this.updateData()
          this.addDialog = false
        } else {
          this.showMsg('warning', '请正确输入')
          return false
        }
      })
    },
    handleCommand (row, command) {
      console.log('data', row, command)
      this.timingId = row.id
      if (command === 'a') {
        homeAjax.startControllers({
          id: row.id
        })
          .then(res => {
            this.showMsg('success', '打铃成功 ^_^')
            this.initController()
          })
          .catch(err => {
            console.log(err)
          })
      } else {
        this.timingFormVisible = true
      }
    },
    confirmClick () {
      if (!this.timingTemp.time) {
        this.showMsg('warning', '请设置时间！')
        return false
      }
      homeAjax.timingControllers({
        id: this.timingId,
        time: translateTime(this.timingTemp.time).hm
      })
        .then(res => {
          this.showMsg('success', '打铃成功 ^_^')
          this.initController()
          this.timingFormVisible = false
        })
        .catch(err => {
          console.log(err)
        })
    },
    refreshClick () {
      this.iniData()
    },
    resetForm () {
      this.$nextTick(() => {
        this.$refs['controlForm'].clearValidate();
      })
    },
    handleAccept (data) {
      homeAjax.closeControllers({
        id: data.id
      })
        .then(res => {
          this.showMsg('success', '断开成功 ^_^')
          this.iniData()
        })
        .catch(err => {
          console.log(err)
        })
    },
    resetFormData () {
      this.controlForm = {
        code: '',
        schedule_id: '',
        name: ''
      }
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
.mainControl-wrapper
  .el-date-editor.el-input
    width 95%
  .el-button--success
    background-color: #009688;
    border-color #009688
    &:hover
      opacity: .8;
      color: #fff;
  .add-dialog
    .el-dialog__title
      color #fff
    .el-dialog__header
      background-color: #20222A
    .el-dialog__footer
      .el-button
        height: 38px;
        line-height: 38px;
        padding: 0 18px;
        font-size: 14px;
        background-color: #009688;
        &:hover
          opacity .8
    .el-select
      width 100%
</style>
<style lang='stylus' scoped>
.mainControl-wrapper
  position: relative;
  min-height 100%
  height auto
  margin: 0 auto;
  box-sizing border-box
  .page-wrapper
    margin-top 20px
</style>
