import React from 'react';

import { Link } from './Link';
import './ul.css';


interface UlProps {
  primary?: boolean;
  titles: string[];
  actives?: boolean[];
  backgroundColor?: string;
  onMenu1: () => void;
  onMenu2: () => void;
  onMenu3: () => void;
  onMenu4: () => void;
}

export const Ul = ({ 
  primary = false,
  titles,
  actives = [false, false, false, false],
  backgroundColor,
  onMenu1,
  onMenu2,
  onMenu3,
  onMenu4,
}: UlProps) => {
  const mode = primary ? 'storybook-ul--primary' : 'storybook-ul--secondary';
  return(
    <ul className={['storybook-ul' , mode].join(' ')}>
      {
        <>
          <Link
            label={titles[0]}
            backgroundColor={backgroundColor}
            onClick={onMenu1}
            primary={primary}
            active={actives[0]}
          />
          <Link
            label={titles[1]}
            backgroundColor={backgroundColor}
            onClick={onMenu2}
            primary={primary}
            active={actives[1]}
          />
          <Link
            label={titles[2]}
            backgroundColor={backgroundColor}
            onClick={onMenu3}
            primary={primary}
            active={actives[2]}
          />
          <Link
            label={titles[3]}
            backgroundColor={backgroundColor}
            onClick={onMenu4}
            primary={primary}
            active={actives[3]}
          />
        </>
      }
    </ul>
  );
}

