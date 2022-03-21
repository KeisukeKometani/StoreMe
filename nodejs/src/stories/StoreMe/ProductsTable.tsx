import * as React from 'react';
import { DataGrid, GridColDef, GridValueGetterParams } from '@mui/x-data-grid';

const columns: GridColDef[] = [
  { field: 'id', headerName: 'ID' },
  { field: 'productCode', headerName: 'Product Code' },
  { field: 'name', headerName: 'Name' },
  { field: 'price', headerName: 'Price' },
  { field: 'quantity', headerName: 'Quantity' },
];

const rows = [
  { id: 1, productCode: 'Snow', name: 'Jon', price: 35, quantity: 1 },
  { id: 2, productCode: 'Lannister', name: 'Cersei', price: 42, quantity: 2 },
  { id: 3, productCode: 'Lannister', name: 'Jaime', price: 45, quantity: 3 },
  { id: 4, productCode: 'Stark', name: 'Arya', price: 16, quantity: 4 },
  { id: 5, productCode: 'Targaryen', name: 'Daenerys', price: 1, quantity: 5 },
  { id: 6, productCode: 'Melisandre', name: 'Ken', price: 150, quantity: 6 },
  { id: 7, productCode: 'Clifford', name: 'Ferrara', price: 44, quantity: 7 },
  { id: 8, productCode: 'Frances', name: 'Rossini', price: 36, quantity: 8 },
  { id: 9, productCode: 'Roxie', name: 'Harvey', price: 65, quantity: 9 },
  { id: 10, productCode: 'Snow', name: 'Jon', price: 35, quantity: 1 },
  { id: 11, productCode: 'Snow', name: 'Jon', price: 35, quantity: 1 },
  { id: 12, productCode: 'Lannister', name: 'Cersei', price: 42, quantity: 2 },
  { id: 13, productCode: 'Lannister', name: 'Jaime', price: 45, quantity: 3 },
  { id: 14, productCode: 'Stark', name: 'Arya', price: 16, quantity: 4 },
  { id: 15, productCode: 'Targaryen', name: 'Daenerys', price: 1, quantity: 5 },
  { id: 16, productCode: 'Melisandre', name: 'Ken', price: 150, quantity: 6 },
  { id: 17, productCode: 'Clifford', name: 'Ferrara', price: 44, quantity: 7 },
  { id: 18, productCode: 'Frances', name: 'Rossini', price: 36, quantity: 8 },
  { id: 19, productCode: 'Roxie', name: 'Harvey', price: 65, quantity: 9 },
];

interface ProductsTableProps {
  pageSize: number;
  rowsPerPageOptions: number[];
  checkboxSelection: boolean;
}

export const ProductsTable = ({
  pageSize,
  rowsPerPageOptions,
  checkboxSelection,
}: ProductsTableProps) => {
  return(
    <div style={{ height: 800, width: '100%' }}>
      <DataGrid
        columns={columns}
        rows={rows}
        pageSize={pageSize}
        rowsPerPageOptions={rowsPerPageOptions}
        checkboxSelection={checkboxSelection}
      />
    </div>
    
  );
};