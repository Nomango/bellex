<template>
  <div
    class="home-wrapper"
    v-loading="ajaxLoading"
    element-loading-text="拼命加载中"
    element-loading-spinner="el-icon-loading"
    element-loading-background="rgba(0, 0, 0, 0.2)">
    <bell-header />
    <bell-menu />
    <bell-body>
      <router-view slot="bellcontent" />
    </bell-body>
  </div>
</template>

<script>
import bellHeader from './components/bellHeader'
import bellMenu from './components/bellMenu'
import bellBody from './components/bellContent'
import { mapState } from 'vuex'
import homeAjax from '@/api/home.js'
import { setSession } from '@/utils/storage.js'
export default {
  name: 'home',
  components: {
    bellMenu,
    bellHeader,
    bellBody
  },
  data () {
    return {
    };
  },
  computed: {
    ...mapState(['ajaxLoading'])
  },
  created() {
    homeAjax.getLogin().then(res => {
      console.log(res)
      setSession('is_login', res.is_login)
    }).catch(err => {
      console.log(err)
    })
  },
  mounted () {

  },
  methods: {

  }
};
</script>

<style lang="stylus">
.home-wrapper .bell-side
  width: 220px;
  top: 0;
  z-index: 1001;
</style>
<style lang="stylus" scoped>
.home-wrapper
  width 100%
  height 100%
</style>
