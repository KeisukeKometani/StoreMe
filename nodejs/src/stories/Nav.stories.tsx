import React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';

import { Nav } from './Nav';

export default {
  title: 'Example/Nav',
  component: Nav,
} as ComponentMeta<typeof Nav>;

const Template: ComponentStory<typeof Nav> = (args) => <Nav {...args} />;

export const MainNav = Template.bind({});
MainNav.args = {
  actives: [false, false, false, false],
  col: 12,
};