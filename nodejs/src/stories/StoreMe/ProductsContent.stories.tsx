import * as React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';
import { ProductsContent } from './ProductsContent';


export default {
  title: 'StoreMe/ProductsContent',
  component: ProductsContent,
} as ComponentMeta<typeof ProductsContent>;


const Template: ComponentStory<typeof ProductsContent> = (args) => <ProductsContent {...args} />;


export const Default = Template.bind({});
Default.args = {
};