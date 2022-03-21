import * as React from 'react';
import MuiListItemText from '@mui/material/ListItemText';

interface ListItemTextProps {
  primary: string;
  secondary?: string;
}

export const ListItemText = ({
  primary,
  secondary,
}: ListItemTextProps) => {
  return (
    <MuiListItemText primary={primary} secondary={secondary} />
  );
};