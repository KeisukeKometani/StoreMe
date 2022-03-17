import React from 'react';

import { Ul } from './Ul';
import './nav.css';


interface NavProps {
  actives: boolean[];
  onMenu1: () => void;
  onMenu2: () => void;
  onMenu3: () => void;
  onMenu4: () => void;
}

export const Nav = ({ 
  actives = [false, false, false, false],
  onMenu1,
  onMenu2,
  onMenu3,
  onMenu4,
}: NavProps) => (
  <nav>
    <div className={'storybook-nav'}>
      {
        <>
          <Ul
            backgroundColor="#212529"
            actives={actives}
            onMenu1={onMenu1}
            onMenu2={onMenu2}
            onMenu3={onMenu3}
            onMenu4={onMenu4}
            primary
            size="large"
            titles={[
              '商品一覧',
              '売上分析',
              'マニュアル',
              'お問い合わせ'
            ]}
          />
        </>
      }
    </div>
  </nav>
);
