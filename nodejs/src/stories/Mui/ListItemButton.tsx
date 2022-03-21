import * as React from 'react';
import MuiListItemButton from '@mui/material/ListItemButton';

import { ListItemIcon } from './ListItemIcon';
import { ListItemText } from './ListItemText';

interface ListItemButtonProps {
  component: React.ComponentType<MuiListItemButton.Props>;
  href?: string;
  iconImageName?: string;
  label: string;
}

export const ListItemButton = ({
  component: Component,
  href,
  iconImageName,
  label
}: ListItemButtonProps) => {
  return (
    <MuiListItemButton 
      component={ Component }
      href={ href }
    >
      <ListItemIcon iconImageName={ iconImageName } />
      <ListItemText primary={ label } />
    </MuiListItemButton>
  );
}