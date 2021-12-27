import React from 'react';
import { Comm } from 'api/Comm'
import { AlgoGetMessage } from 'api/message/admin'

const Landing = () => {
    return (
        <div>
            <h1 className={'h1'}>
                Landing
            </h1>
            <button onClick={() => Comm.get(AlgoGetMessage())}>
                AlgoGet
            </button>
        </div>
    );
};

export default Landing;