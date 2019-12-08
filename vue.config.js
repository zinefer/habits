const path = require("path");

module.exports = {
  lintOnSave: false,
  outputDir: path.resolve(__dirname, "web/dist"),
  configureWebpack: {
    entry: "./web/src/main.js",
    resolve: {
      alias: {
        "@": path.resolve(__dirname, "web/src")
      }
    }
  }
};
