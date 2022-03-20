import React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';
import { MUIButton } from './MUIButton';

export default {
  title: 'StoreMe/MUIButton',
  component: MUIButton,
} as ComponentMeta<typeof MUIButton>;

const Template: ComponentStory<typeof MUIButton> = (args) => <MUIButton {...args} />;

export const Default = Template.bind({});
Default.args = {
  label: 'Button',
};