import React from 'react';
import './list.css';

interface ListProps {
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
  }

  export const List = ({
    primary = false,
    size = 'medium',
    backgroundColor,
    label,
    ...props
  }: ListProps) => {
    const mode = primary ? 'storybook-list--primary' : 'storybook-list--secondary';
    return (
      <li
        className={['storybook-list' ,`storybook-list--${size}`, mode].join(' ')}
        style={{ backgroundColor }}
        {...props}
      >
        {label}
      </li>
    );
  };