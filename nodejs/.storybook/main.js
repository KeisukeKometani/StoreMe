const HardSourceWebpackPlugin = require('hard-source-webpack-plugin')

module.exports = {
  "stories": [
    "../src/**/*.stories.(js|mdx)",
    "../src/**/*.stories.@(js|jsx|ts|tsx)"
  ],
  "addons": [
    "@storybook/addon-links",
    "@storybook/addon-essentials",
    "@storybook/addon-interactions"
  ],
  "plugins": [
    new HardSourceWebpackPlugin()
  ],
  "framework": "@storybook/react"
}