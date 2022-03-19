import merge from 'deepmerge';
import path from 'path';
import { createBasicConfig } from '@open-wc/building-rollup';
import serve from 'rollup-plugin-serve';
import livereload from 'rollup-plugin-livereload';
import typescript from 'rollup-plugin-typescript2';
import postcss from 'rollup-plugin-postcss';
import commonjs from '@rollup/plugin-commonjs';
import replace from '@rollup/plugin-replace';
import sass from 'rollup-plugin-sass';

// TODO: プロダクション環境の場合の対応
const baseConfig = createBasicConfig({ developmentMode: true });

export default merge(baseConfig, {
  //input: './out-tsc/src/index.js',
  input: './src/index.tsx',
  output: {
    dir: 'dist',
    format: 'iife',
    sourceMap: true,
  },
  plugins: [
    serve({
      open: true,
      verbose: true,
      contentBase: ['', 'public'],
      host: '0.0.0.0',
      port: 8082
    }),
    livereload({
      watch: "dist",
      port: 8090
    }),
    typescript(),
    postcss({ plugins: [] }),
    commonjs(),
    replace({
      // TODO: プロダクション環境の場合の対応
      'process.env.NODE_ENV': JSON.stringify('development'),
    }),
    sass({
      output: 'dist/style.css',
    })
  ]
});