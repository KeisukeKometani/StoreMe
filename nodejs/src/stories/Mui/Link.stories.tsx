import * as React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';

import { Link } from './Link';

export default {
  title: 'Mui/Link',
  component: Link,
} as ComponentMeta<typeof Link>;

const Template: ComponentStory<typeof Link> = (args) => <Link {...args} />;

export const Primary = Template.bind({});
Primary.args = {
  label: 'Link',
  color: 'primary',
};


export const Secondary = Template.bind({});
Secondary.args = {
  label: 'Link',
};


export const H1 = Template.bind({});
H1.args = {
  label: 'Link',
  variant: 'h1',
};


export const UnderLineHover = Template.bind({});
UnderLineHover.args = {
  label: 'Link',
  underline: 'hover',
};