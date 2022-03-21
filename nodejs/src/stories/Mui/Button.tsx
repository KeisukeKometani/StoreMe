import * as React from 'react';
import MuiButton from '@mui/material/Button';
import SendIcon from '@mui/icons-material/Send';
import ViewInArIcon from '@mui/icons-material/ViewInAr';

type variantTypes = 'text' | 'outlined' | 'contained';
type colorTypes = 'inherit' | 'primary' | 'secondary' | 'success' | 'info' | 'warning' | 'error';
type sizeTypes = 'small' | 'medium' | 'large';
type IconImageTypes = 'NULL' | 'send' | 'viewInAr';

interface ButtonProps {
  label: string;
  variant?: variantTypes;
  disabled?: boolean;
  href?: string;
  disableElevation?: boolean;
  color?: colorTypes;
  size?: sizeTypes;
  startIconImageType?: IconImageTypes;
  endIconImageType?: IconImageTypes;
  onclick?: () => void;
}

export const Button = ({
  label,
  variant = 'contained',
  disabled = false,
  href,
  disableElevation,
  color = 'inherit',
  size = 'medium',
  startIconImageType,
  endIconImageType,
  ...props
}: ButtonProps) => {
  const startIcon = iconImage(startIconImageType);
  const endIcon = iconImage(endIconImageType);

  return (
    <MuiButton
      variant={variant}
      disabled={disabled}
      href={href}
      disableElevation={disableElevation}
      color={color}
      size={size}
      startIcon={startIcon}
      endIcon={endIcon}
      {...props}
    >
      {label}
    </MuiButton>
  );
}

function iconImage(startIconImageType: IconImageTypes): JSX.Element {
  switch (startIconImageType) {
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