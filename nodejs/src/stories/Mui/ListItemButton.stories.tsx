import * as React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';
import { ListItemButton } from './ListItemButton';

export default {
  title: 'Mui/ListItemButton',
  component: ListItemButton,
} as ComponentMeta<typeof ListItemButton>;

const Template: ComponentStory<typeof ListItemButton> = (args) => <ListItemButton {...args} />;

export const ProductAll = Template.bind({});
ProductAll.args = {
  component: 'a',
  href: 'https://www.google.com',
  iconImageName: 'viewInAr',
  label: '商品一覧',
};

export const SalesAnalysis = Template.bind({});
SalesAnalysis.args = {
  component: 'a',
  href: 'https://www.google.com',
  iconImageName: 'equalizer',
  label: '売上分析',
};

export const Manual = Template.bind({});
Manual.args = {
  component: 'a',
  href: 'https://www.google.com',
  iconImageName: 'menuBook',
  label: 'マニュアル',
};

export const Help = Template.bind({});
Help.args = {
  component: 'a',
  href: 'https://www.google.com',
  iconImageName: 'help',
  label: 'お問い合わせ',
};