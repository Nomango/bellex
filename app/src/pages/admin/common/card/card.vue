<template>
  <div class="bell-card">
    <div class="card-header">
      <slot name="card-header" />
    </div>
    <div class="card-content">
      <slot name="card-content" />
    </div>
  </div>
</template>
<script>
export default {
}
</script>
<style lang='stylus' scoped>
.bell-card
  height 100%
  border-radius: 2px;
  background-color: #fff;
  box-shadow: 0 1px 2px 0 rgba(0,0,0,.05);
  .card-header
    height: 42px;
    line-height: 42px;
    padding: 10px 15px;
    border-bottom: 1px solid #f6f6f6;
  .card-content
    position: relative;
    padding: 10px 15px;
    line-height: 24px;
</style>
