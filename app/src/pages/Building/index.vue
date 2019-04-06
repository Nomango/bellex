<template>
  <div class="build-wrapper">
    <bell-card>
      <div class="card-header-content" slot="card-header">
        <el-button @click="addTableClick" type="success">添加</el-button>
      </div>
      <div class="card-content-body" slot="card-content">
          <el-table
            :data="buildingData"
            stripe
            border
            style="width: 100%">
          <el-table-column
            prop="name"
            label="名称"
            width="" />
          <el-table-column
            prop="date"
            label="时间"
            width="" />
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
            :dialogVisible="buildingDialog"
            @confirm="handleConfirm"
            @cancel="handleCancel" />
        </div>
        <div class="add-dialog">
          <add-dialog
            width="40%"
            :dialogVisible="addDialog"
            @cancel="handleAddCancel"
            @confirm="handleAddConfirm">
            <div class="content" slot="content">
              <el-form ref="buildingForm" :model="buildingForm" label-width="80px">
                <el-form-item label="名称: ">
                  <el-input v-model="buildingForm.name"></el-input>
                </el-form-item>
                <el-form-item label="时间表: ">
                  <div class="time-wrapper">
                    <el-tag
                      v-if="buildingForm.timeTags.length"
                      :key="tag"
                      v-for="tag in buildingForm.timeTags"
                      closable
                      :disable-transitions="true"
                      @close="handleCloseTag(tag)">
                      {{tag}}
                    </el-tag>
                  </div>
                  <el-time-picker
                    v-if="timeVisible"
                    class="input-new-tag"
                    v-model="timeValue"
                    format="HH:mm"
                    :picker-options="{
                      selectableRange: '00:00:00 - 23:59:59'
                    }"
                    placeholder="请选择时间"
                    ref="saveTagInput"
                    @keyup.enter.native="handlePickerConfirm"
                    @blur="handlePickerConfirm" />
                  <el-button
                    v-else
                    class="button-new-tag"
                    icon="iconfont icon-add"
                    size="small" @click="showTimepicker">添加时间</el-button>
                </el-form-item>
                <el-form-item class="footer-item">
                  <el-button type="primary" @click="onSubmit">确定</el-button>
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
import bellCard from '@/pages/common/card/card'
import bellPagination from '@/pages/common/pagination/pagination'
import bellDialog from '@/pages/common/dialog/dialog'
import addDialog from '@/pages/common/dialog/addDialog'
import { translateTime } from '@/utils/tools.js'
export default {
  components: {
    bellCard,
    bellPagination,
    bellDialog,
    addDialog
  },
  data () {
    return {
      timeVal: new Date(2019, 4, 2, 18, 40),
      timeVisible: false,
      timeValue: '',
      buildingDialog: false,
      addDialog: false,
      currentPage: 1,
      totalPage: 100,
      pageSizes: {
        sizeArr: [5, 8, 10, 20],
        size: 5
      },
      buildingForm: {
        name: '',
        timeTags: []
      },
      buildingData: [{
        date: '2016-05-02',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄'
      }, {
        date: '2016-05-04',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1517 弄'
      }, {
        date: '2016-05-01',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1519 弄'
      }, {
        date: '2016-05-03',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1516 弄'
      }]
    }
  },
  methods: {
    handleCheck (val) {
      console.log('handleCheck', val)
    },
    handleSizeChange (val) {
      console.log('handleSizeChange', val)
    },
    handleCurrentChange (val) {
      console.log('handleCurrentChange', val)
    },
    handleDelete (val) {
      this.buildingDialog = true
      console.log('delete', val)
    },
    handleCancel (val) {
      this.buildingDialog = val
    },
    handleConfirm (val) {
      this.buildingDialog = val
      console.log('handleConfirm', val)
    },
    handleAddCancel (val) {
      this.addDialog = val
    },
    handleAddConfirm (val) {
      this.addDialog = val
    },
    addTableClick () {
      this.addDialog = true
    },
    onSubmit () {
      this.addDialog = false
      console.log('kkk')
    },
    handleCloseTag (tag) {
      this.buildingForm.timeTags.splice(this.buildingForm.timeTags.indexOf(tag), 1);
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
        if (this.buildingForm.timeTags.indexOf(newTimeVal.hms) < 0) {
          this.buildingForm.timeTags.push(newTimeVal.hms)
        } else {
          this.$message({
            message: '时间已存在',
            type: 'warning'
          })
        }
      }
      this.timeVisible = false
      this.timeValue = ''
    }
  }
}
</script>
<style lang="stylus">
.build-wrapper
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
    margin-left: 10px;
    vertical-align: bottom;
    .el-input__inner
      height: 30px;
      line-height: 30px;
</style>
<style lang='stylus' scoped>
.build-wrapper
  position: relative;
  height 100%
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
