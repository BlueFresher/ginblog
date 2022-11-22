const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  // 关闭eslint校验
  lintOnSave: false ,
  publicPath: '/admin/',
  outputDir: 'dist',
  assetsDir:'static',
})

// module.exports = {
//   lintOnSave: false ,
//   publicPath: '/admin/',
//   outputDir: 'dist',
//   assetsDir:'static',
// }
