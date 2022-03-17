import React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';
import { within, userEvent } from '@storybook/testing-library';
import { Header } from './Header';

export default {
  title: 'Example/Header',
  component: Header,
  // More on argTypes: https://storybook.js.org/docs/react/api/argtypes
  argTypes: {
    backgroundColor: { control: 'color' },
  },
} as ComponentMeta<typeof Header>;

const Template: ComponentStory<typeof Header> = (args) => <Header {...args} />;

export const LoggedOut = Template.bind({});
LoggedOut.args = {
  backgroundColor: '#212529',
};

export const LoggedIn = Template.bind({});
LoggedIn.args = {
  backgroundColor: '#212529',
};

LoggedIn.play = async ({ canvasElement }) => {
  const canvas = within(canvasElement);
  const loginButton = await canvas.getByRole('button', { name: /新規登録/i });
  await userEvent.click(loginButton);
};