<template>
  <div class="institution-wrapper">
    <bell-card>
      <div class="card-header-content" slot="card-header">
        <el-button @click="addTableClick" type="success">添加</el-button>
      </div>
      <div class="card-content-body" slot="card-content">
          <el-table
            :data="institutionLists"
            stripe
            border
            style="width: 100%">
          <el-table-column
            prop="name"
            label="名称"
            align="center"
            width="" />
          <el-table-column
            prop="create_time"
            label="创建时间"
            align="center"
            width="" />
          <el-table-column
            prop="address"
            label="操作"
            align="center"
            width="200">
            <template slot-scope="scope">
              <el-button @click="handleEdit(scope.row)" icon="el-icon-edit" size="mini">编辑</el-button>
              <el-button @click="handleDelete(scope.row)" icon="el-icon-delete" type="danger" size="mini">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
        <div class="page-wrapper">
          <bell-pagination
            v-show="totalPage > 0"
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
              <el-form ref="timeForm" :model="timeForm" label-width="80px">
                <el-form-item label="名称: ">
                  <el-input v-model="timeForm.name"></el-input>
                </el-form-item>
                <el-form-item class="footer-item">
                  <el-button type="primary" @click="clickStatus==='create'? createData() : updateData()">确定</el-button>
                </el-form-item>
              </el-form>
            </div>
          </add-dialog>
        </div>
      </div>
    </bell-card>
  </div>
</template>
<script>
import bellCard from 'common/card/card'
import bellPagination from 'common/pagination/pagination'
import addDialog from 'common/dialog/addDialog'
import { translateTime } from '@/utils/tools.js'
import institutionAjax from '@/api/institution.js'
export default {
  components: {
    bellCard,
    bellPagination,
    addDialog
  },
  data () {
    return {
      timeVisible: false,
      timeValue: '',
      addDialog: false,
      currentPage: 1,
      totalPage: 0,
      pageSizes: {
        sizeArr: [5, 8, 10, 20],
        size: 5
      },
      clickType: {
        update: '编辑',
        create: '添加'
      },
      clickStatus: null,
      timeForm: {
        name: ''
      },
      timeId: null,
      institutionLists: []
    }
  },
  created () {
    this.getTimeData()
  },
  methods: {
    getTimeData () {
      institutionAjax.getInstitutionList({
        page: this.currentPage,
        limit: this.pageSizes.size
      })
        .then(res => {
          console.log('getInstitutionList', res)
          res = res.data
          this.totalPage = res.total
          this.handleResData(res.data)
        })
        .catch(err => {
          console.log(err)
        })
    },
    handleResData (res) {
      this.institutionLists = res.map(item => {
        return {
          create_time: translateTime(item.create_time).all,
          id: item.id,
          name: item.name
        }
      })
    },
    handleEdit (row) {
      console.log('handleCheck', row)
      this.timeForm = Object.assign({}, row)
      this.clickStatus = 'update'
      this.addDialog = true
      this.timeId = row.id
    },
    handleSizeChange (val) {
      console.log('handleSizeChange', val)
      this.pageSizes.size = val
      this.getTimeData()
    },
    handleCurrentChange (val) {
      console.log('handleCurrentChange', val)
      this.currentPage = val
      this.getTimeData()
    },
    handleDelete (val) {
      console.log('delete', val)
      this.$confirm('此操作将永久删除, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        institutionAjax.delInstitutionList({
          id: val.id
        })
          .then(res => {
            console.log('del', res)
            this.showMsg('success', '删除成功!')
            this.getTimeData()
          })
          .catch(err => {
            console.log(err)
          })
      }).catch(() => {
        this.showMsg('info', '已取消删除')
      })
    },
    handleAddCancel (val) {
      this.addDialog = val
    },
    handleAddConfirm (val) {
      this.addDialog = val
    },
    addTableClick () {
      this.clickStatus = 'create'
      this.addDialog = true
      this.resetForm()
    },
    updateData () {
      let timeForm = this.timeForm
      if (timeForm.name) {
        institutionAjax.putInstitutionList({
          id: this.timeId,
          name: timeForm.name
        })
          .then(res => {
            this.addDialog = false
            this.showMsg('success', '编辑成功')
            this.getTimeData()
          })
          .catch(err => {
            console.log(err)
          })
      }
    },
    createData () {
      let timeForm = this.timeForm
      if (timeForm.name) {
        institutionAjax.addInstitutionList({
          name: timeForm.name
        })
          .then(res => {
            this.addDialog = false
            this.showMsg('success', '添加成功')
            this.getTimeData()
          })
          .catch(err => {
            console.log(err)
          })
      } else {
        this.showMsg('warning', '请输入名称')
      }
    },
    handleCloseTag (tag) {
      this.timeForm.list.splice(this.timeForm.list.indexOf(tag), 1);
    },
    showTimepicker () {
      this.timeVisible = true
      this.$nextTick(_ => {
        this.$refs.saveTagInput.focus()
      });
    },
    handlePickerConfirm () {
      let timeValue = this.timeValue
      if (timeValue) {
        let newTimeVal = translateTime(timeValue)
        if (this.timeForm.list.indexOf(newTimeVal.hm) < 0) {
          this.timeForm.list.push(newTimeVal.hm)
        } else {
          this.showMsg('warning', '时间已存在')
        }
      }
      this.timeVisible = false
      this.timeValue = ''
    },
    resetForm () {
      this.timeForm = {
        name: '',
        list: []
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
.institution-wrapper
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
      display none
    .footer-item
      .el-button
        height: 38px;
        line-height: 38px;
        padding: 0 18px;
        font-size: 14px;
        background-color: #009688;
        &:hover
          opacity .8
    .el-button--small
      padding: 0 9.2px;
  .el-tag + .el-tag
    margin-left: 10px;
  .button-new-tag
    height: 32px;
    line-height: 30px;
    padding-top: 0;
    padding-bottom: 0;
    .icon-add
      display: inline-block;
      vertical-align: bottom;
      margin-right: 5px;
  .input-new-tag
    width: 120px;
    vertical-align: bottom;
    .el-input__inner
      height: 30px;
      line-height: 30px;
</style>
<style lang='stylus' scoped>
.institution-wrapper
  position: relative;
  min-height 100%
  height auto
  margin: 0 auto;
  box-sizing border-box
  .page-wrapper
    margin-top 20px
  .add-dialog .content
    .time-add
      display inline-block
      cursor pointer
      vertical-align middle
      .icon-add
        color #009688
        font-size 20px
        &:hover
          opacity .8
</style>
