import * as React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';
import { Typography } from './Typography';
export default {
    title: 'Mui/Typography',
    component: Typography,
} as ComponentMeta<typeof Typography>;


const Template: ComponentStory<typeof Typography> = (args) => <Typography {...args} />;


export const Primary = Template.bind({});
Primary.args = {
    label: 'Primary',
    color: 'primary',
};