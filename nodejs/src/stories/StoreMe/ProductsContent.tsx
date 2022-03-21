import * as React from 'react';
import Container from '@mui/material/Container';
import Typography from '@mui/material/Typography';
import { ProductsTable } from './ProductsTable';

interface ProductsContentProps {
}

export const ProductsContent = ({

}: ProductsContentProps) => {
  return (
    <Container>
      <Typography variant="h4" component="div" gutterBottom>
        { '商品一覧' }
      </Typography>
      <ProductsTable
        checkboxSelection
        pageSize={10}
        rowsPerPageOptions={[
          10
        ]}
      />
    </Container>
  );
};