module.exports = {
  lintOnSave: false,
  configureWebpack: {
    entry: "./web/src/main.js",
    resolve: {
      alias: {
        "@": require("path").resolve(__dirname, "web/src")
      }
    }
  }
};
