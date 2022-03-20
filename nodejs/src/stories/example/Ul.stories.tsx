import React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';

import { Ul } from './Ul';

export default {
  title: 'Example/Ul',
  component: Ul,
  argTypes: {
    backgroundColor: { control: 'color' },
  },
} as ComponentMeta<typeof Ul>;

const Template: ComponentStory<typeof Ul> = (args) => <Ul {...args} />;

export const Primary = Template.bind({});
Primary.args = {
  primary: true,
  titles: ['Menu', 'Menu2', 'Menu3', 'Menu4'],
};

export const Secondary = Template.bind({});
Secondary.args = {
  titles: ['Menu', 'Menu2', 'Menu3', 'Menu4'],
};

export const Large = Template.bind({});
Large.args = {
  size: 'large',
  titles: ['Menu', 'Menu2', 'Menu3', 'Menu4'],
};

export const Active = Template.bind({});
Active.args = {
  titles: ['Menu', 'Menu2', 'Menu3', 'Menu4'],
  actives: [true, false, false, false],
};

