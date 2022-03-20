import React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';

import { List } from './List';

export default {
    title: 'Example/List',
    component: List,
    // More on argTypes: https://storybook.js.org/docs/react/api/argtypes
    argTypes: {
        backgroundColor: { control: 'color' },
    },
} as ComponentMeta<typeof List>;

const Template: ComponentStory<typeof List> = (args) => <List {...args} />;

export const Primary = Template.bind({});
Primary.args = {
    primary: true,
    label: 'List',
};

export const Secondary = Template.bind({});
Secondary.args = {
    label: 'List',
};

