import * as React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';
import { MUITabs } from './MUITabs';

export default {
  title: 'StoreMe/MUITabs',
  component: MUITabs,
} as ComponentMeta<typeof MUITabs>;


const Template: ComponentStory<typeof MUITabs> = (args) => <MUITabs {...args} />;


export const MainMenu = Template.bind({});
MainMenu.args = {
  labels: ['商品一覧', '売上分析', 'マニュアル', 'お問い合わせ'],
  icons: ['viewInAr', 'equalizer', 'menuBook', 'help'],
};