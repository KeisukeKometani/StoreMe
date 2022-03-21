import * as React from 'react';
import Link from '@mui/material/Link';

type variantTypes = 'body1' | 'body2' | 'button' | 'caption' | 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6' | 'inherit' | 'overline' | 'subtitle1' | 'subtitle2';
type colorTypes = 'inherit' | 'primary' | 'secondary';
type underlineTypes = 'none' | 'hover' | 'always';

interface MUILinkProps {
  label: string;
  variant?: variantTypes;
  href?: string;
  color?: colorTypes;
  underline?: underlineTypes;
  onclick?: () => void;
}

export const MUILink = ({
  label,
  variant = 'body1',
  href,
  color = 'inherit',
  underline = 'hover',
  ...props
}: MUILinkProps) => {
  return (
    <Link
      variant={variant}
      href={href}
      color={color}
      underline={underline}
      {...props}
    >
      {label}
    </Link>
  );
}