import * as React from 'react';
import ListItemIcon from '@mui/material/ListItemIcon';
import SendIcon from '@mui/icons-material/Send';
import ViewInArIcon from '@mui/icons-material/ViewInAr';
import EqualizerIcon from '@mui/icons-material/Equalizer';
import MenuBookIcon from '@mui/icons-material/MenuBook';
import HelpIcon from '@mui/icons-material/Help';

interface MUIListItemIconProps {
  iconImageName: 'string';
}

export const MUIListItemIcon = ({
  iconImageName,
}: MUIListItemIconProps) => {
  const icon = iconImage(iconImageName);

  return (
    <ListItemIcon>
      {icon}
    </ListItemIcon>
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