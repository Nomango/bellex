<template>
  <div class="userManage-wrapper">
    <bell-card>
      <div class="card-header-content" slot="card-header">
        <el-button @click="addClick" type="success">添加</el-button>
      </div>
      <div class="card-content-body" slot="card-content">
        <el-table
            :data="subUserData"
            stripe
            border
            style="width: 100%">
          <el-table-column
            prop="name"
            label="名称"
            width="" />
          <el-table-column
            prop="account"
            label="账号"
            width="" />
          <el-table-column
            prop="password"
            label="密码"
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
            :total="totalPage"
            :currentPage="currentPage"
            @sizeChange="handleSizeChange"
            @currentChange="handleCurrentChange" />
        </div>
        <div class="normal-dialog">
          <bell-dialog
            width="24%"
            :dialogVisible="subUserDialog"
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
              <el-form ref="subUserForm" :model="subUserForm" label-width="60px">
                <el-form-item label="名称: ">
                  <el-input v-model="subUserForm.name" placeholder="请输入名称" />
                </el-form-item>
                <el-form-item label="账号: ">
                  <el-input v-model="subUserForm.account" placeholder="请输入账号" />
                </el-form-item>
                <el-form-item label="密码: ">
                  <el-input v-model="subUserForm.password" placeholder="请输入密码"/>
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
export default {
  components: {
    bellCard,
    bellPagination,
    bellDialog,
    addDialog
  },
  data () {
    return {
      totalPage: 20,
      currentPage: 1,
      subUserDialog: false,
      addDialog: false,
      subUserForm: {
        name: '',
        account: '',
        password: ''
      },
      subUserData: [{
        name: '张三',
        account: 'leoGoGo',
        password: '123456'
      }, {
        name: '李四',
        account: 'leoGoGo',
        password: '123456'
      }, {
        name: '刘六',
        account: 'leoGoGo',
        password: '123456'
      }]
    }
  },
  mounted () {
  },
  methods: {
    addClick () {
      this.addDialog = true
    },
    handleEdit (data) {
      this.addDialog = true
    },
    handleDelete (data) {
      this.subUserDialog = true
    },
    handleSizeChange (val) {
      console.log('handleSizeChange', val)
    },
    handleCurrentChange (val) {
      console.log('handleCurrentChange', val)
    },
    handleConfirm (val) {
      this.subUserDialog = val
    },
    handleCancel (val) {
      this.subUserDialog = val
    },
    handleAddConfirm (val) {
      this.addDialog = val
    },
    handleAddCancel (val) {
      this.addDialog = val
    }
  }
}
</script>
<style lang="stylus">
.userManage-wrapper
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
      .el-button
        height: 38px;
        line-height: 38px;
        padding: 0 18px;
        font-size: 14px;
        background-color: #009688;
        &:hover
          opacity .8
</style>
<style lang='stylus' scoped>
.userManage-wrapper
  position: relative;
  height 100%
  margin: 0 auto;
  box-sizing border-box
  .page-wrapper
    margin-top 20px
</style>
