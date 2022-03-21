import * as React from 'react';
import MuiListItemIcon from '@mui/material/ListItemIcon';
import SendIcon from '@mui/icons-material/Send';
import ViewInArIcon from '@mui/icons-material/ViewInAr';
import EqualizerIcon from '@mui/icons-material/Equalizer';
import MenuBookIcon from '@mui/icons-material/MenuBook';
import HelpIcon from '@mui/icons-material/Help';

interface ListItemIconProps {
  iconImageName: 'string';
}

export const ListItemIcon = ({
  iconImageName,
}: ListItemIconProps) => {
  const icon = iconImage(iconImageName);

  return (
    <MuiListItemIcon>
      {icon}
    </MuiListItemIcon>
  );
}

function iconImage(iconImageName: string): JSX.Element {
  switch (iconImageName) {
    case 'NULL':
      return;
    case 'send':
      return <SendIcon />;
    case 'viewInAr':
      return <ViewInArIcon />;
    case 'equalizer':
      return <EqualizerIcon />;
    case 'menuBook':
      return <MenuBookIcon />;
    case 'help':
      return <HelpIcon />;
    default:
      return '';
  }
}