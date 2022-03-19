import React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';

import { MButton } from './MButton';

export default {
  title: 'Example/MButton',
  component: MButton,
} as ComponentMeta<typeof MButton>;

const Template: ComponentStory<typeof MButton> = (args) => <MButton {...args} />;

export const Primary = Template.bind({});