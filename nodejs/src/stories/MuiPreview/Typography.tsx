import * as React from 'react';
import Box from '@mui/material/Box';
import MuiTypography from '@mui/material/Typography';

interface TypographyProps {
  color?: 'inherit' | 'primary' | 'secondary' | 'textPrimary' | 'textSecondary' | 'error';
}

export const Typography = ({
  color = 'inherit',
}: TypographyProps) => {
  return (
    <Box>
      <MuiTypography variant={ 'h1' } component={ 'h1' } color={color}>h1</MuiTypography>
      <MuiTypography variant={ 'h2' } component={ 'h2' } color={color}>h2</MuiTypography>
      <MuiTypography variant={ 'h3' } component={ 'h3' } color={color}>h3</MuiTypography>
      <MuiTypography variant={ 'h4' } component={ 'h4' } color={color}>h4</MuiTypography>
      <MuiTypography variant={ 'h5' } component={ 'h5' } color={color}>h5</MuiTypography>
      <MuiTypography variant={ 'h6' } component={ 'h6' } color={color}>h6</MuiTypography>
      <MuiTypography variant={ 'subtitle1' } component={ 'subtitle1' } color={color}>subtitle1</MuiTypography>
      <MuiTypography variant={ 'subtitle2' } component={ 'subtitle2' } color={color}>subtitle2</MuiTypography>
      <MuiTypography variant={ 'body1' } component={ 'body1' } color={color}>body1</MuiTypography>
      <MuiTypography variant={ 'body2' } component={ 'body2' } color={color}>body2</MuiTypography>
      <MuiTypography variant={ 'caption' } component={ 'caption' } color={color}>caption</MuiTypography>
      <MuiTypography variant={ 'button' } component={ 'button' } color={color}>button</MuiTypography>
      <MuiTypography variant={ 'overline' } component={ 'overline' } color={color}>overline</MuiTypography>
    </Box>
    
  );
};