'use strict'
const utils = require('./utils')
const config = require('../config')

module.exports = {
  loaders: utils.cssLoaders({
    sourceMap: config.cssSourceMap,
    extract: true
  }),
  cssSourceMap: config.cssSourceMap,
  cacheBusting: config.cacheBusting,
  transformToRequire: {
    video: ['src', 'poster'],
    source: 'src',
    img: 'src',
    image: 'xlink:href'
  }
}
