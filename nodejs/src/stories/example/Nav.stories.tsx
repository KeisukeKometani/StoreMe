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
};

export const LoggedIn = Template.bind({});
LoggedIn.args = {
  user: {
    name: 'Jane Doe',
  },
  actives: [false, false, false, false]
};

export const LoggedOut = Template.bind({});
LoggedOut.args = {
  actives: [false, false, false, false]
};