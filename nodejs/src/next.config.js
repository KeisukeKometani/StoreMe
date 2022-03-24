function useEsbuildLoader(config, options) {
  const jsLoader = config.module.rules.find(
    rule => rule.test && rule.test.test(".js")
  );

  if (jsLoader) {
    jsLoader.use.loader = "esbuild-loader";
    jsLoader.use.options = options;
  }
}

module.exports = withTypescript({
  webpack(config, options) {
    config.plugins.push(new webpack.EnvironmentPlugin(process.env));

    config.plugins.push(
      new webpack.ProvidePlugin({
        React: "react"
      })
    );

    useEsbuildLoader(config, {
      loader: "tsx",
      target: "es2017"
    });

    if (options.isServer) {
      config.plugins.push(
        new ForkTsCheckerWebpackPlugin({
          tsconfig: path.resolve("./tsconfig.json"),
          memoryLimit: 1024
        })
      );
    }
    if (process.env.NODE_ENV && process.env.NODE_ENV === "production") {
      config.optimization.minimize = true;
      config.optimization.minimizer = [
        new TerserPlugin({
          terserOptions: {
            compress: { drop_console: true }
          },
          extractComments: "all"
        })
      ];
    }
    return config;
  }
})