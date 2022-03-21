import * as React from 'react';
import ListItemButton from '@mui/material/ListItemButton';

import { MUIListItemIcon } from './MUIListItemIcon';
import { MUIListItemText } from './MUIListItemText';

interface MUIListItemButtonProps {
  component: React.ComponentType<ListItemButton.Props>;
  href?: string;
  iconImageName?: string;
  label: string;
}

export const MUIListItemButton = ({
  component: Component,
  href,
  iconImageName,
  label
}: MUIListItemButtonProps) => {
  return (
    <ListItemButton 
      component={ Component }
      href={ href }
    >
      <MUIListItemIcon iconImageName={ iconImageName } />
      <MUIListItemText primary={ label } />
    </ListItemButton>
  );
}