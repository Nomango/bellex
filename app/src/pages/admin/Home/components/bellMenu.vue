<template>
  <div class="bell-side bell-side-menu" :class="{'isHideMenu': isCollapse}">
    <div class="bell-logo">
      <span>{{userInfo.institution.name}}</span>
    </div>
    <el-menu
      :default-active="defaultActive"
      :collapse="isCollapse"
      :unique-opened="true"
      class="el-menu-vertical-demo"
      @select="handleSelect"
      background-color="#20222A"
      text-color="rgba(255,255,255,.7)"
      active-text-color="#ffffff">
      <el-submenu
        v-for="item of roleMenu"
        :key="item.id"
        :index="item.id">
        <template slot="title">
          <i class="iconfont" :class="item.icon" />
          <span>{{item.name}}</span>
        </template>
        <el-menu-item
          v-if="item.children && item.children.length"
          v-for="(subItem,index) of item.children"
          :key="index"
          :index="subItem.index">
          {{subItem.name}}
        </el-menu-item>
      </el-submenu>
    </el-menu>
  </div>
</template>
<script>
import { mapState } from 'vuex'
export default {
  computed: {
    ...mapState(['isCollapse']),
    ...mapState(['roles']),
    ...mapState(['userInfo'])
  },
  data () {
    return {
      defaultActive: this.$route.path || '/home/mainControl',
      roleMenu: null,
      menuList: [{
        id: '1',
        icon: 'icon-home',
        name: '主页',
        children: [{
          index: '/home/mainControl',
          name: '主控机管理'
        }, {
          index: '/home/building',
          name: '时间表'
        }]
      }, {
        id: '2',
        icon: 'icon-userManage',
        name: '高级管理',
        children: [{
          index: '/home/userManage',
          name: '用户管理'
        }, {
          index: '/home/institution',
          name: '机构管理'
        }]
      }, {
        id: '3',
        icon: 'icon-shezhi',
        name: '设置',
        children: [{
          index: '/home/profile',
          name: '修改资料'
        }, {
          index: '/home/set',
          name: '修改密码'
        }]
      }]
    }
  },
  created () {
    this.handleRoleMenu()
  },
  methods: {
    handleRoleMenu () {
      if (this.roles === 0) {
        let newMenuArr = Object.assign([], this.menuList)
        newMenuArr.splice(1, 1)
        this.roleMenu = newMenuArr
      } else {
        this.roleMenu = this.menuList
      }
    },
    handleSelect (index, indexPath) {
      if (index) {
        this.$router.push(index)
      }
    }
  },
  watch: {
    $route (newVal, oVal) {
      this.defaultActive = newVal.path
    },
    'roles' (newVal, oVal) {
      if (newVal) {
        this.roleMenu = this.menuList
      } else {
        let newMenuArr = Object.assign([], this.menuList)
        this.roleMenu = newMenuArr.splice(1, 1)
      }
    }
  }
}
</script>
<style lang="stylus">
.bell-side
  &.isHideMenu
    .el-menu--inline, .el-submenu__title span,
    .el-submenu__icon-arrow
      display none
  .el-menu
    left 1px
  .el-menu-item.is-active
    background-color #009688!important
  .el-submenu .el-menu-item, .el-menu-item, .el-submenu__title
    height: 40px;
    line-height: 40px;
  .el-menu-vertical-demo
    margin-top 50px
  .el-submenu__title
    .iconfont
      margin-right 5px
      vertical-align middle
      line-height 44px
</style>
<style lang='stylus' scoped>
.bell-side
  position: fixed;
  left: 0;
  top: 0;
  bottom: 0;
  width: 220px;
  top: 0;
  z-index: 1001;
  overflow-x: hidden;
  transition: all .3s;
  &.isHideMenu
    width 64px
    .bell-logo
      width: 60px;
      background: #fff;
      span
        display none
  &.bell-side-menu
    color: #fff;
    box-shadow: 1px 0 2px 0 rgba(0,0,0,.05);
    background-color #20222A
  .bell-logo
    position: fixed;
    left: 0;
    top: 0;
    z-index: 1002;
    width: 218.1px;
    height: 49px;
    line-height 49px
    padding: 0 15px;
    text-align center
    box-sizing: border-box;
    font-weight: 300;
    transition: all .3s;
    color: rgba(255,255,255,.8);
    background-color #20222A
    background-repeat: no-repeat;
    background-position: center center;
</style>
