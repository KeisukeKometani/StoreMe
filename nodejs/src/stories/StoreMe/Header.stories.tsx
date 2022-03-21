import * as React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';
import { Header } from './Header';


export default {
  title: 'StoreMe/Header',
  component: Header,
} as ComponentMeta<typeof Header>;


const Template: ComponentStory<typeof Header> = (args) => <Header {...args} />;


export const Default = Template.bind({});
Default.args = {
  position: 'static',
};