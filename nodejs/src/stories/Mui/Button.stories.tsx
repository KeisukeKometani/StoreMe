import * as React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';
import { Button } from './Button';

export default {
  title: 'Mui/Button',
  component: Button,
} as ComponentMeta<typeof Button>;

const Template: ComponentStory<typeof Button> = (args) => <Button {...args} />;

export const Primary = Template.bind({});
Primary.args = {
  color: 'primary',
  label: 'Button',
};

export const Secondary = Template.bind({});
Secondary.args = {
  label: 'Button',
};

export const Small = Template.bind({});
Small.args = {
  label: 'Button',
  size: 'small',
};

export const Large = Template.bind({});
Large.args = {
  label: 'Button',
  size: 'large',
};

export const Text = Template.bind({});
Text.args = {
  variant: 'text',
  label: 'Button',
  color: 'primary',
};

export const Outlined = Template.bind({});
Outlined.args = {
  variant: 'outlined',
  label: 'Button',
  color: 'primary',
};

export const Disabled = Template.bind({});
Disabled.args = {
  label: 'Button',
  disabled: true,
  color: 'primary',
};

// boxのshadowなどが消えます。
export const DisableElevation = Template.bind({});
DisableElevation.args = {
  label: 'Button',
  disableElevation: true,
  color: 'primary',
};

export const startIcon = Template.bind({});
startIcon.args = {
  label: 'Button',
  startIconImageType: 'send',
  color: 'primary',
};

export const endIcon = Template.bind({});
endIcon.args = {
  label: 'Button',
  endIconImageType: 'viewInAr',
  color: 'primary',
};