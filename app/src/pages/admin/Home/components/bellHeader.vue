<template>
  <div class="bell-header" :class="{'isCollapse': isCollapse}">
    <div class="bell-nav bell-layout-left">
      <span class="bell-nav-item" @click="handleChangeCollapse">
        <i class="iconfont" :class="isShousuo" />
      </span>
    </div>
    <div class="bell-nav bell-layout-right">
      <el-dropdown @visible-change="handleVisible">
        <span class="el-dropdown-link">
          我的<i class="el-icon--right" :class="dropIcon" />
        </span>
        <el-dropdown-menu slot="dropdown">
          <!-- <el-dropdown-item icon="">基本资料</el-dropdown-item> -->
          <el-dropdown-item icon="">
            <router-link to="/home/set">修改密码</router-link>
          </el-dropdown-item>
          <el-dropdown-item icon="">
            <router-link @click.native="handleLogOut">退出</router-link>
          </el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </div>
  </div>
</template>
<script>
import { mapMutations, mapState } from 'vuex'
import loginAjax from '@/api/login.js'
export default {
  data () {
    return {
      dropIcon: 'el-icon-caret-bottom'
    }
  },
  computed: {
    ...mapState(['isCollapse']),
    isShousuo () {
      return this.isCollapse ? 'icon-zhankai' : 'icon-shousuo'
    }
  },
  methods: {
    ...mapMutations(['changeCollapse']),
    handleLogOut() {
      loginAjax.logout()
        .then(res => {
          this.$message({
            message: '登出成功!',
            type: 'success'
          })
          setTimeout(() => {
            this.$router.go(0)
          }, 500)
        })
        .catch(err => {
          console.log(err)
        })
    },
    handleVisible (val) {
      if (val) {
        this.dropIcon = 'el-icon-caret-top'
      } else {
        this.dropIcon = 'el-icon-caret-bottom'
      }
    },
    handleChangeCollapse () {
      this.changeCollapse()
    }
  }
}
</script>
<style lang="stylus">
.el-dropdown-menu__item
  text-align center
</style>
<style lang='stylus' scoped>
.bell-header
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 50px;
  box-sizing: border-box;
  background-color: #fff;
  box-shadow 0 1px 2px 0 rgba(0,0,0,.05);
  z-index: 1000;
  &.isCollapse
    .bell-nav.bell-layout-left
      left 64px
  .bell-nav
    height 100%
    position absolute
    top: 0;
    transition all .3s
    .bell-nav-item
      display inline-block
      font-size 20px
      cursor pointer
    &.bell-layout-left
      left 220px
      padding 0 10px
      line-height 50px
    &.bell-layout-right
      right 50px
      .el-dropdown-link
        cursor pointer
        line-height 50px
        outline none
</style>
