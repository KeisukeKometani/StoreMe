import * as React from 'react';
import MuiLink from '@mui/material/Link';

type variantTypes = 'body1' | 'body2' | 'button' | 'caption' | 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6' | 'inherit' | 'overline' | 'subtitle1' | 'subtitle2';
type colorTypes = 'inherit' | 'primary' | 'secondary';
type underlineTypes = 'none' | 'hover' | 'always';

interface LinkProps {
  label: string;
  variant?: variantTypes;
  href?: string;
  color?: colorTypes;
  underline?: underlineTypes;
  onclick?: () => void;
}

export const Link = ({
  label,
  variant = 'body1',
  href,
  color = 'inherit',
  underline = 'hover',
  ...props
}: LinkProps) => {
  return (
    <MuiLink
      variant={variant}
      href={href}
      color={color}
      underline={underline}
      {...props}
    >
      {label}
    </MuiLink>
  );
}