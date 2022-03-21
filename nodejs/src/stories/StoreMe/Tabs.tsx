import * as React from 'react';
import Box from '@mui/material/Box';
import TabContext from '@mui/lab/TabContext';
import TabList from '@mui/lab/TabList';
import TabPanel from '@mui/lab/TabPanel';
import Tab from '@mui/material/Tab';
import { ListItemIcon } from '../Mui/ListItemIcon'

interface TabsProps {
  labels: string[];
  icons: string[];
  tabPanelContents: React.ReactNode[];
}

export const Tabs = ({ 
  labels,
  icons,
  tabPanelContents,
 }: TabsProps) => {
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
              icon={<ListItemIcon iconImageName={ icons[i] } />}
              iconPosition="start"
              value={`${i + 1}`}
            />
          ))}
        </TabList>
      </Box>
      {labels.map((label, i) => (
        <TabPanel value={`${i + 1}`}>
          {tabPanelContents[i]}
        </TabPanel>
      ))}
    </TabContext>
  );
};