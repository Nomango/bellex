<template>
  <el-pagination
    @size-change="sizeChange"
    @current-change="currentChange"
    :current-page="currentPage"
    :page-sizes="pageSizes.sizeArr"
    :page-size="pageSizes.size"
    layout="total, sizes, prev, pager, next, jumper"
    :total="total">
  </el-pagination>
</template>
<script>
export default {
  props: {
    currentPage: Number,
    pageSizes: {
      type: Object,
      default: () => {
        return {
          sizeArr: [5, 8, 10, 20],
          size: 5
        }
      }
    },
    total: Number
  },
  data () {
    return {
    }
  },
  methods: {
    sizeChange (val) {
      this.$emit('sizeChange', val)
    },
    currentChange (val) {
      this.$emit('currentChange', val)
    }
  }
}
</script>
<style lang='stylus' scoped>

</style>
