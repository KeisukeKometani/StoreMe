import * as React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';
import { Tabs } from './Tabs';
import { ProductsContent } from './ProductsContent';

export default {
  title: 'StoreMe/Tabs',
  component: Tabs,
} as ComponentMeta<typeof Tabs>;


const Template: ComponentStory<typeof Tabs> = (args) => <Tabs {...args} />;


export const MainMenu = Template.bind({});
MainMenu.args = {
  labels: ['商品一覧', '売上分析', 'マニュアル', 'お問い合わせ'],
  icons: ['viewInAr', 'equalizer', 'menuBook', 'help'],
  tabPanelContents: [<ProductsContent />, <ProductsContent />, <ProductsContent />, <ProductsContent />],
};