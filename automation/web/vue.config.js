module.exports = {
  devServer: {
    headers: {
      'Access-Control-Allow-Origin': '*',
    },
    proxy: {
      '/api': {
        target: 'http://10.10.10.10:8888',
        ws: false,
        changeOrigin: true,
        pathRewrite: {
          // '^/api': '',
        },
      },
    },
    port: 4032,
    host: 'localhost',
  },
  // 静态资源
  publicPath: process.env.NODE_ENV === 'development' ? '/' : './',
  outputDir: 'dist',
  assetsDir: 'static/',
  indexPath: 'index.html',
  filenameHashing: true,
  // 关闭默认的vue-cli-service lint功能，下方plugins中手动配置，以处理eslint9时编译报错的问题
  lintOnSave: true,
  chainWebpack: (config) => {
    config.resolve.symlinks(true);
  },
  configureWebpack: {
    output: {
      //资源打包路径
      library: 'vueApp',
      libraryTarget: 'umd',
    },
  },
};
