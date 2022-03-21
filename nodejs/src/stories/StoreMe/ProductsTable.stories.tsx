import * as React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';
import { ProductsTable } from './ProductsTable';


export default {
  title: 'StoreMe/ProductsTable',
  component: ProductsTable,
} as ComponentMeta<typeof ProductsTable>;


const Template: ComponentStory<typeof ProductsTable> = (args) => <ProductsTable {...args} />;


export const Default = Template.bind({});
Default.args = {
  pageSize: 5,
  rowsPerPageOptions: [10],
  checkboxSelection: true
};
