const path = require('path');

module.exports = {
  mode: "development",
  devtool: "inline-source-map",
  entry: "./src/app.tsx",
  output: {
    publicPath: path.join(__dirname, '..', 'api', 'public'),
    path: path.join(__dirname, '..', 'api', 'public'),
    filename: "index.js",
  },
  resolve: {
    extensions: [".ts", ".tsx", ".js", ".jsx"]
  },
  module: {
    rules: [
      { test: /\.tsx?$/, loader: "ts-loader" },
      { test: /\.scss$/, use: ["style-loader", "css-loader", "sass-loader"] }
    ]
  }
};