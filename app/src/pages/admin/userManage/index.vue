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
            prop="username"
            label="名称"
            align="center"
            width="" />
          <el-table-column
            prop="email"
            label="邮箱"
            align="center"
            width="" />
          <el-table-column
            label="角色"
            align="center"
            width="">
            <template slot-scope="{row}">
              <el-tag :type="row.role | tagsFilter">{{row.role | roleFilter}}</el-tag>
            </template>
          </el-table-column>
           <el-table-column
            label="机构"
            align="center"
            width="">
            <template slot-scope="{row}">
              <span>{{row.insititution.name}}</span>
            </template>
          </el-table-column>
          <el-table-column
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
                <el-form-item label="名称: " prop="username">
                  <el-input v-model="subUserForm.username" placeholder="请输入名称" />
                </el-form-item>
                <el-form-item label="邮箱: " prop="email">
                  <el-input v-model="subUserForm.email" placeholder="请输入账号" />
                </el-form-item>
                <el-form-item label="机构: " prop="institution">
                  <el-select v-model="subUserForm.institution" placeholder="请选择活动区域">
                    <el-option
                      v-for="item of institutionOption"
                      :key="item.value"
                      :label="item.label"
                      :value="item.value" />
                  </el-select>
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
  filters: {
    roleFilter (status) {
      const statusMap = {
        0: '普通管理员',
        1: '一级管理员',
        2: '超级管理员'
      }
      return statusMap[status]
    },
    tagsFilter (status) {
      const statusMap = {
        0: 'info',
        1: 'success',
        2: 'primary'
      }
      return statusMap[status]
    }
  },
  data () {
    return {
      totalPage: 0,
      currentPage: 1,
      pageSizes: 5,
      subUserDialog: false,
      addDialog: false,
      subUserForm: {
        username: '',
        email: '',
        institution: ''
      },
      subUserData: [],
      clickType: null,
      institutionOption: [],
      rules: {
        username: [{ required: true, message: '请输入名称', trigger: 'blur' }],
        email: [{ required: true, message: '请输入账号', trigger: 'blur' }],
        institution: [{ required: true, message: '请选择机构', trigger: 'change' }]
      }
    }
  },
  created () {
    this.getUserList()
    this.getInstitution()
  },
  methods: {
    getInstitution () {
      userAjax.getInstitution({
        page: 0,
        limit: 0
      })
        .then(res => {
          res = res.data
          this.handleInstitutionData(res)
        })
        .catch(err => {
          console.log(err)
        })
    },
    handleInstitutionData (res) {
      res = res.data
      let result = res.map(item => {
        return {
          label: item.name,
          value: item.id
        }
      })
      this.institutionOption = result
      console.log('getInstitution', this.institutionOption)
    },
    getUserList () {
      userAjax.getUserList({
        page: this.currentPage,
        limit: this.pageSizes
      })
        .then(res => {
          res = res.data
          console.log('res', res)
          this.subUserData = res.data
          this.totalPage = res.total
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
      this.subUserForm.institution = data.insititution.name
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
      this.pageSizes = val
      this.getUserList()
    },
    handleCurrentChange (val) {
      console.log('handleCurrentChange', val)
      this.currentPage = val
      this.getUserList()
    },
    handleAddConfirm (val) {
      this.$refs['subUserForm'].validate((valid) => {
        if (valid) {
          this.clickType === 'create' ? this.createData(val) : this.updateData(val)
          console.log('submit!')
        } else {
          console.log('error submit!!')
          return false
        }
      });
    },
    createData (val) {
      let subUserForm = this.subUserForm
      userAjax.addUser({
        username: subUserForm.username,
        email: subUserForm.email,
        institution_id: subUserForm.institution
      })
        .then(res => {
          this.showMsg('success', '添加成功!')
        })
        .catch(err => {
          console.log(err)
        })
      this.addDialog = val
    },
    updateData (val) {
      let subUserForm = this.subUserForm
      userAjax.putUser({
        username: subUserForm.username,
        email: subUserForm.email,
        institution_id: subUserForm.institution
      })
        .then(res => {
          this.showMsg('success', '编辑成功!')
        })
        .catch(err => {
          console.log(err)
        })
      this.addDialog = val
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
    .el-select
      width 100%
</style>
<style lang='stylus' scoped>
.userManage-wrapper
  position: relative;
  min-height 100%
  height auto
  margin: 0 auto;
  box-sizing border-box
  .page-wrapper
    margin-top 20px
</style>
