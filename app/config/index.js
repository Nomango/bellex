const path = require('path')

module.exports = {
  // Paths
  assetsRoot: path.resolve(__dirname, '../../server'),
  assetsSubDirectory: 'static',
  assetsPublicPath: '/',

  index: path.resolve(__dirname, '../../server/views/index.html'),

  /**
   * Source Maps
   */

  cssSourceMap: false,
  // https://webpack.js.org/configuration/devtool/#production
  devtool: '#source-map',

  // Gzip off by default as many popular static hosts such as
  // Surge or Netlify already gzip all static assets for you.
  // Before setting to `true`, make sure to:
  // npm install --save-dev compression-webpack-plugin
  productionGzip: false,
  productionGzipExtensions: ['js', 'css'],

  // If you have problems debugging vue-files in devtools,
  // set this to false - it *may* help
  // https://vue-loader.vuejs.org/en/options.html#cachebusting
  cacheBusting: true,
}
