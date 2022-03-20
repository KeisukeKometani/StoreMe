import React from 'react';
import Button from '@mui/material/Button';
import SendIcon from '@mui/icons-material/Send';

type variantTypes = 'text' | 'outlined' | 'contained';
type colorTypes = 'inherit' | 'primary' |'secondary' | 'success' | 'error' | 'info' | 'warning';
type sizeTypes = 'small' | 'medium' | 'large';

interface MUIButtonProps {
  label: string;
  variant?: variantTypes;
  disabled?: boolean;
  href?: string;
  disableElevation?: boolean;
  color?: colorTypes;
  size?: sizeTypes;
  onclick?: () => void;
}

export const MUIButton = ({
  label,
  variant = 'contained',
  disabled = false,
  href,
  disableElevation,
  color = 'inherit',
  size = 'medium',
  ...props
}: MUIButtonProps) => {
  return(
    <Button
      variant={variant}
      disabled={disabled}
      href={href}
      disableElevation={disableElevation}
      color={color}
      size={size}
      startIcon={<SendIcon />}
      {...props}
    >
      {label}
    </Button>
  );
}