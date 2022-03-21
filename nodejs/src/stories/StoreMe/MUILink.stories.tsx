import * as React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';

import { MUILink } from './MUILink';

export default {
  title: 'StoreMe/MUILink',
  component: MUILink,
} as ComponentMeta<typeof MUILink>;

const Template: ComponentStory<typeof MUILink> = (args) => <MUILink {...args} />;

export const Default = Template.bind({});
Default.args = {
  label: 'Link',
};