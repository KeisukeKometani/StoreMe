import * as React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';
import { IconButton } from './IconButton';


export default {
  title: 'Mui/IconButton',
  component: IconButton,
} as ComponentMeta<typeof IconButton>;


const Template: ComponentStory<typeof IconButton> = (args) => <IconButton {...args} />;


export const Menu = Template.bind({});
Menu.args = {
  ImageType: 'menu',
};
