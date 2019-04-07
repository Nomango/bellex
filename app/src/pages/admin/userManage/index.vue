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
        <div class="add-dialog">
          <add-dialog
            width="40%"
            :dialogVisible="addDialog"
            @cancel="handleAddCancel"
            @confirm="handleAddConfirm">
            <div class="content" slot="content">
              <el-form ref="subUserForm" :model="subUserForm" :rules="rules" label-width="60px">
                <el-form-item label="名称: " prop="name">
                  <el-input v-model="subUserForm.name" placeholder="请输入名称" />
                </el-form-item>
                <el-form-item label="账号: " prop="account">
                  <el-input v-model="subUserForm.account" placeholder="请输入账号" />
                </el-form-item>
                <el-form-item label="密码: " prop="password">
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
import bellCard from 'common/card/card'
import bellPagination from 'common/pagination/pagination'
import addDialog from 'common/dialog/addDialog'
import userAjax from '@/api/user.js'
export default {
  components: {
    bellCard,
    bellPagination,
    addDialog
  },
  data () {
    return {
      totalPage: 0,
      currentPage: 1,
      subUserDialog: false,
      addDialog: false,
      subUserForm: {
        name: '',
        account: '',
        password: ''
      },
      subUserData: [],
      clickType: null,
      rules: {
        name: [{ required: true, message: '请输入名称', trigger: 'blur' }],
        account: [{ required: true, message: '请输入账号', trigger: 'blur' }],
        password: [{ required: true, message: '请输入密码', trigger: 'blur' }]
      }
    }
  },
  created () {
    this.getUserList()
  },
  methods: {
    getUserList () {
      userAjax.getUserList()
        .then(res => {
          console.log('user', res)
          if (res.code === 0) {
            res = res.data
            this.subUserData = res.userList
            this.totalPage = res.total
          }
        })
        .catch(err => {
          console.log(err)
        })
    },
    addClick () {
      this.resetFormData()
      this.addDialog = true
      this.clickType = 'create'
      this.resetForm()
    },
    handleEdit (data) {
      this.subUserForm = Object.assign({}, data)
      this.clickType = 'update'
      this.addDialog = true
      this.resetForm()
    },
    handleDelete (data) {
      this.$confirm('此操作将永久删除, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        userAjax.delUser({
          id: data.id
        })
          .then(res => {
            if (res.code === 0) {
              this.showMsg('success', '删除成功!')
            }
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
    },
    handleCurrentChange (val) {
      console.log('handleCurrentChange', val)
    },
    handleAddConfirm (val) {
      this.$refs['subUserForm'].validate((valid) => {
        if (valid) {
          this.addDialog = val
          console.log('submit!')
        } else {
          console.log('error submit!!')
          return false
        }
      });
    },
    handleAddCancel (val) {
      this.addDialog = val
    },
    resetFormData () {
      this.subUserForm = {
        name: '',
        account: '',
        password: ''
      }
    },
    resetForm () {
      this.$nextTick(() => {
        this.$refs['subUserForm'].clearValidate();
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
.userManage-wrapper
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
