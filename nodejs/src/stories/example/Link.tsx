import * as React from 'react';
import './link.css';

interface LinkProps {
  /**
   * Is this the principal call to action on the page?
   */
  primary?: boolean;
  /**
   * What background color to use
   */
  backgroundColor?: string;
  /**
   * How large should the button be?
   */
  size?: 'small' | 'medium' | 'large';
  /**
   * Button contents
   */
  label: string;
  /**
   * Active page
   */
  active?: boolean;
  /**
   * Optional click handler
   */
  onClick?: () => void;
}

export const Link = ({
  primary = false,
  active = false,
  size = 'medium',
  backgroundColor,
  label,
  ...props
}: LinkProps) => {
  const mode = primary ? 'storybook-link--primary' : 'storybook-link--secondary';
  const activeMode = active ? 'storybook-link--active' : '';
  return (
    <a
      className={['storybook-link', `storybook-link--${size}`, mode, activeMode].join(' ')}
      style={{ backgroundColor }}
      {...props}
    >
      {label}
    </a>
  );
};