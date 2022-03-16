import React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';

import { Link } from './Link';

export default {
    title: 'Example/Link',
    component: Link,
    // More on argTypes: https://storybook.js.org/docs/react/api/argtypes
    argTypes: {
        backgroundColor: { control: 'color' },
    },
} as ComponentMeta<typeof Link>;

const Template: ComponentStory<typeof Link> = (args) => <Link {...args} />;

export const Primary = Template.bind({});
Primary.args = {
    primary: true,
    label: 'Link',
};

export const Secondary = Template.bind({});
Secondary.args = {
  label: 'Link',
};

export const Active = Template.bind({});
Active.args = {
  active: true,
  label: 'Link',
};

  