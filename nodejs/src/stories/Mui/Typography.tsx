import * as React from 'react';
import MuiTypography from '@mui/material/Typography';

interface TypographyProps {
  label: string;
  variant?: 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6' | 'subtitle1' | 'subtitle2' | 'body1' | 'body2' | 'caption' | 'button' | 'overline';
  color?: 'inherit' | 'primary' | 'secondary' | 'textPrimary' | 'textSecondary' | 'error';
}

export const Typography = ({
  label,
  variant = 'h1',
  color = 'inherit',
}: TypographyProps) => {
  return (
    <MuiTypography
      variant={ variant }
      component={ variant }
      color={color}
    >
      { label}
    </MuiTypography>
  );
};