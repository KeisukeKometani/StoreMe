import * as React from 'react';
import MuiIconButton from '@mui/material/IconButton';
import MenuIcon from '@mui/icons-material/Menu';
import SendIcon from '@mui/icons-material/Send';
import ViewInArIcon from '@mui/icons-material/ViewInAr';

type colorTypes = 'inherit' | 'primary' | 'secondary' | 'success' | 'info' | 'warning' | 'error';
type sizeTypes = 'small' | 'medium' | 'large';
type IconImageTypes = 'NULL' | 'menu' | 'send' | 'viewInAr';

interface IconButtonProps {
  disabled?: boolean;
  disableRipple?: boolean;
  disableFocusRipple?: boolean;
  href?: string;
  color?: colorTypes;
  size?: sizeTypes;
  edge?: 'start' | 'end';
  ImageType?: IconImageTypes;
  onclick?: () => void;
}

export const IconButton = ({
  disabled = false,
  disableRipple = false,
  disableFocusRipple = false,
  href,
  color = 'inherit',
  size = 'medium',
  edge = 'start',
  ImageType,
  ...props
}: IconButtonProps) => {
  const Icon = iconImage(ImageType);
  
  return (
    <MuiIconButton
      disabled={disabled}
      disableRipple={disableRipple}
      disableFocusRipple={disableFocusRipple}
      href={href}
      color={color}
      size={size}
      edge={edge}
      {...props}
    >
      {Icon}
    </MuiIconButton>
  );
}

function iconImage(IconImageType: IconImageTypes): JSX.Element {
  switch (IconImageType) {
    case 'NULL':
      return;
    case 'menu':
      return <MenuIcon />;
    case 'send':
      return <SendIcon />;
    case 'viewInAr':
      return <ViewInArIcon />;
    default:
      return '';
  }
}