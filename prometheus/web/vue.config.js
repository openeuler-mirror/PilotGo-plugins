module.exports = {
  devServer: {
    headers: {
      'Access-Control-Allow-Origin': '*',
    },
    proxy: {
      '/': {
        target: 'http://localhost:8082',
        ws: false,
        changeOrigin: true,
      }
    },
    port: 8082,
    host: 'localhost',
  },

  // 静态资源
  publicPath: "./",
  outputDir: "dist",//"../server/resource/dist",
  assetsDir: "static",
  indexPath: "index.html",
  filenameHashing: true,

  chainWebpack: (config) => {
    config.resolve.symlinks(true);
  },
  configureWebpack: {
    output: {
      //资源打包路径
      library: 'vueApp',
      libraryTarget: 'umd'
    }
  }
}
