import * as React from 'react';
import ReactDom from 'react-dom';
import { Page } from './stories/example/Page';
import './test.sass';
import './style/reset.sass';

ReactDom.render(<Page />, document.querySelector('#root'));

