const webpack = require('webpack')

module.exports = [{
  entry: './src/js/app.js',
  output: {
    path: __dirname + '/app/static/js/',
    filename: 'bundle.js'
  },
  module: {
    rules: [{
      test: /\.tag$/,
      enforce: 'pre',
      exclude: /node_modules/,
      use: [{
        loader: 'riot-tag-loader',
        options: {
          hot: false,
          debug: true
        }
      }]
    }, {
      test: /\.js|\.tag$/,
      enforce: 'post',
      exclude: /node_modules/,
      use: [{
        loader: 'babel-loader',
        options: {
          presets: `es2015-riot`
        }
      }]
    }]
  },
  resolve: {
    extensions: ['.js', '.tag']
  },
  plugins: [
    new webpack.ProvidePlugin({ riot: 'riot' })
  ]
}]
