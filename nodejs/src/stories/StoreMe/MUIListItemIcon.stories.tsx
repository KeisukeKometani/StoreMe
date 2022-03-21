import * as React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';

import { MUIListItemIcon } from './MUIListItemIcon';

export default {
  title: 'StoreMe/MUIListItemIcon',
  component: MUIListItemIcon,
} as ComponentMeta<typeof MUIListItemIcon>;

const Template: ComponentStory<typeof MUIListItemIcon> = (args) => <MUIListItemIcon {...args} />;

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