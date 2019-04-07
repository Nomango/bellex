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
      </div>
    </bell-card>
  </div>
</template>
<script>
import bellCard from 'common/card/card'
import bellPagination from 'common/pagination/pagination'
import homeAjax from '@/api/home.js'
export default {
  components: {
    bellCard,
    bellPagination
  },
  data () {
    return {
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
      this.$router.push({
        name: 'tableDetail',
        params: {
          id: val.id
        }
      })
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
  .el-button--success
    background-color: #009688;
    border-color #009688
    &:hover
      opacity: .8;
      color: #fff;
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
