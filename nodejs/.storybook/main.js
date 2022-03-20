const HardSourceWebpackPlugin = require('hard-source-webpack-plugin')

module.exports = {
  "stories": [
    "../src/**/*.stories.(js|mdx)",
    "../src/**/*.stories.@(js|jsx|ts|tsx)"
  ],
  "addons": [
    "@storybook/addon-links",
    "@storybook/addon-essentials",
    "@storybook/addon-interactions",
    'storybook-addon-material-ui5'
  ],
  webpackFinal(config) {
    delete config.resolve.alias['styled-components'];
    delete config.resolve.alias['mui/styled-engine-sc'];
    return config;
  },
  "plugins": [
    new HardSourceWebpackPlugin()
  ],
  "framework": "@storybook/react"
}