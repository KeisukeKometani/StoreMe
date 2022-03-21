import * as React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';

import { ListItemIcon } from './ListItemIcon';

export default {
  title: 'Mui/ListItemIcon',
  component: ListItemIcon,
} as ComponentMeta<typeof ListItemIcon>;

const Template: ComponentStory<typeof ListItemIcon> = (args) => <ListItemIcon {...args} />;

export const Send = Template.bind({});
Send.args = {
  iconImageName: 'send',
};

export const ViewInAr = Template.bind({});
ViewInAr.args = {
  iconImageName: 'viewInAr',
};

export const Equalizer = Template.bind({});
Equalizer.args = {
  iconImageName: 'equalizer',
};

export const MenuBook = Template.bind({});
MenuBook.args = {
  iconImageName: 'menuBook',
};

export const Help = Template.bind({});
Help.args = {
  iconImageName: 'help',
};