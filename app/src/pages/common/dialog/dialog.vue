<template>
  <el-dialog
    title="提示"
    :visible.sync="dialogVisible"
    :modal-append-to-body="false"
    :width="width"
    top="15%"
    :before-close="handleClose">
    <span>确定删除吗</span>
    <span slot="footer" class="dialog-footer">
      <el-button @click="cancelClick">取 消</el-button>
      <el-button type="primary" @click="confirmClick">确 定</el-button>
    </span>
</el-dialog>
</template>
<script>
export default {
  props: {
    width: {
      type: String,
      default: '30%'
    },
    dialogVisible: {
      type: Boolean,
      default: false
    }
  },
  data () {
    return {
      isShowDialog: this.dialogVisible
    }
  },
  methods: {
    confirmClick () {
      this.$emit('confirm', false)
    },
    cancelClick () {
      this.$emit('cancel', false)
    },
    handleClose () {
      this.cancelClick()
    }
  }
}
</script>
<style lang='stylus' scoped>

</style>
