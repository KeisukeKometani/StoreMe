import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import { Typography } from '../Mui/Typography';
import { IconButton } from '../Mui/IconButton';

type positionTypes = 'fixed' | 'absolute' | 'sticky' | 'static' | 'relative';

interface HeaderProps {
  position?: positionTypes;
}


export const Header = ({ position = 'fixed' }: HeaderProps) => {
  return (
    <AppBar position={position}>
      <Toolbar>
        <IconButton
          ImageType="menu"
          edge="start"
        />
        <Typography
          color="textSecondary"
          label="StoreMe"
          variant="h6"
        />
      </Toolbar>
    </AppBar>
  );
};