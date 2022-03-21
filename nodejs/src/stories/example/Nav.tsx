import * as React from 'react';

import { Button } from './Button';
import { Ul } from './Ul';
import './nav.css';

type User = {
  name: string;
};

interface NavProps {
  user?: User;
  actives: boolean[];
  onLogin: () => void;
  onLogout: () => void;
  onCreateAccount: () => void;
  onMenu1: () => void;
  onMenu2: () => void;
  onMenu3: () => void;
  onMenu4: () => void;
}

export const Nav = ({
  user,
  actives = [true, false, false, false],
  onLogin,
  onLogout,
  onCreateAccount,
  onMenu1,
  onMenu2,
  onMenu3,
  onMenu4
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
      <div>
        {user ? (
          <>
            <span className="welcome">
              name: <b>{user.name}</b>!
            </span>
            <Button
              backgroundColor="#C70044"
              size="small"
              label="ログアウト"
              onClick={onLogout}
              primary
              round={5}
            />
          </>
        ) : (
          <>
            <Button
              backgroundColor="transparent"
              size="small"
              label="ログイン"
              onClick={onLogin}
              primary
              round={5}
            />
            <Button
              size="small"
              label="新規登録"
              onClick={onCreateAccount}
              primary
              round={5}
            />
          </>
        )}
      </div>
    </div>
  </nav>
);
