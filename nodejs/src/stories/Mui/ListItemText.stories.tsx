import * as React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';

import { ListItemText } from './ListItemText';

export default {
  title: 'Mui/ListItemText',
  component: ListItemText,
} as ComponentMeta<typeof ListItemText>;

const Template: ComponentStory<typeof ListItemText> = (args) => <ListItemText {...args} />;

export const Primary = Template.bind({});
Primary.args = {
  primary: 'メインテキスト',
};

export const Secondary = Template.bind({});
Secondary.args = {
  primary: 'メインテキスト',
  secondary: 'サブテキスト',
};