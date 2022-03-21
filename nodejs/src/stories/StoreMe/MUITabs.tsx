import * as React from 'react';
import Box from '@mui/material/Box';
import TabContext from '@mui/lab/TabContext';
import TabList from '@mui/lab/TabList';
import TabPanel from '@mui/lab/TabPanel';
import Tab from '@mui/material/Tab';
import { MUIListItemIcon } from './MUIListItemIcon'
import { ProductsContent } from './ProductsContent';

interface MUITabsProps {
  labels: string[];
  icons: string[];
}

export const MUITabs = ({ labels, icons }: MUITabsProps) => {
  const [value, setValue] = React.useState('1');

  const handleChange = (event: React.SyntheticEvent, newValue: string) => {
    setValue(newValue);
  };

  return (
    <TabContext value={value}>
      <Box>
        <TabList onChange={handleChange}>
          {labels.map((label, i) => (
            <Tab
              label={label}
              icon={<MUIListItemIcon iconImageName={ icons[i] } />}
              iconPosition="start"
              value={`${i + 1}`}
            />
          ))}
        </TabList>
      </Box>
      <TabPanel value="1">
        <ProductsContent />
      </TabPanel>
    </TabContext>
  );
};