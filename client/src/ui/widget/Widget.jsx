import React from 'react';
import classes from './Widget.module.css'

const Widget = ({children}) => {
    return (
        <div className={classes.widget}>
            {children}
        </div>
    );
};

export default Widget;