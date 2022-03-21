import * as React from 'react';
import ListItemText from '@mui/material/ListItemText';

interface MUIListItemTextProps {
  primary: string;
  secondary?: string;
}

export const MUIListItemText = ({
  primary,
  secondary,
}: MUIListItemTextProps) => {
  return (
    <ListItemText primary={primary} secondary={secondary} />
  );
};