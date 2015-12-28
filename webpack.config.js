var path = require('path');
var webpack = require('webpack');

var defaultPlugins = [
    new webpack.HotModuleReplacementPlugin()
];

const PATHS = {
    app: path.join(__dirname, 'app'),
    build: path.join(__dirname, 'content')
};

const PROD = JSON.parse(process.env.PROD || "0");

module.exports = {
    entry: PATHS.app,
    output: {
        path: PATHS.build,
        filename: 'bundle.js'
    },
    resolve: {
        extensions: ['', '.js', '.jsx']
    },
    module: {
        loaders: [
        {
            test: /\.css$/,
            loaders: ['style', 'css'],
            include: PATHS.app
        },
        {
            test: /\.jsx?$/,
            loaders: ['babel'],
            include: PATHS.app
        }
        ]
    },
    devServer: {
        historyApiFallback: true,
        hot: true,
        inline: true,
        progress: true,

        // Display only errors to reduce the amount of output.
        stats: 'errors-only',

        // Parse host and port from env so this is easy to customize.
        host: process.env.HOST,
        port: process.env.PORT
    },
    plugins: PROD ? defaultPlugins.concat([ 
        new webpack.optimize.UglifyJsPlugin({ minimize: true }) 
    ]) : defaultPlugins
};
