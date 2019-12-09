const path = require("path");

const CopyWebpackPlugin = require("copy-webpack-plugin");

module.exports = {
  lintOnSave: false,
  outputDir: path.resolve(__dirname, "web/dist"),
  configureWebpack: {
    entry: "./web/src/main.js",
    resolve: {
      alias: {
        "@": path.resolve(__dirname, "web/src")
      }
    },
    plugins: [
      new CopyWebpackPlugin([
        {
          from: path.resolve(__dirname, "web/public"),
          to: path.resolve(__dirname, "web/dist")
        }
      ])
    ]
  },
  chainWebpack: config => {
    config.plugin("html").tap(args => {
      args[0].template = path.resolve(__dirname, "web/public/index.html");
      return args;
    });
  },
  transpileDependencies: ["vuetify"]
};
