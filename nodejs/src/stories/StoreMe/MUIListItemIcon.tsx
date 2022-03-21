import * as React from 'react';
import ListItemIcon from '@mui/material/ListItemIcon';
import SendIcon from '@mui/icons-material/Send';
import ViewInArIcon from '@mui/icons-material/ViewInAr';

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
    default:
      return '';
  }
}