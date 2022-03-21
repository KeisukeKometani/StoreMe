import * as React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';

import { MUIListItemText } from './MUIListItemText';

export default {
  title: 'StoreMe/MUIListItemText',
  component: MUIListItemText,
} as ComponentMeta<typeof MUIListItemText>;

const Template: ComponentStory<typeof MUIListItemText> = (args) => <MUIListItemText {...args} />;

export const Primary = Template.bind({});
Primary.args = {
  primary: 'メインテキスト',
};

export const Secondary = Template.bind({});
Secondary.args = {
  primary: 'メインテキスト',
  secondary: 'サブテキスト',
};