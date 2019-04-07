<template>
  <el-dialog
    :title="title"
    :visible.sync="dialogVisible"
    :modal-append-to-body="false"
    :width="width"
    top="15%"
    :before-close="handleClose">
    <slot name="content"></slot>
    <span slot="footer" class="dialog-footer">
      <el-button type="primary" @click="confirmClick">确 定</el-button>
    </span>
</el-dialog>
</template>
<script>
export default {
  props: {
    title: {
      type: String,
      default: '添加'
    },
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
    handleClose () {
      this.$emit('cancel', false)
    }
  }
}
</script>
<style lang='stylus' scoped>

</style>
