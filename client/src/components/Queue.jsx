import { AuthContext } from 'context';
import React, { useContext } from 'react';

const Queue = () => {
    const {fContext} = useContext(AuthContext)

    return (
        <h1>Queue</h1>
    );
};

export default Queue;