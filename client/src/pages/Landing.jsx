import React, { useContext } from 'react';
import { Comm } from 'api/Comm'
import { AlgoGetMessage } from 'api/message/admin'
import { AuthContext } from 'context';

const Landing = () => {
    const {fContext} = useContext(AuthContext)
    return (
        <div>
            <h1 className={'h1'}>
                Landing
            </h1>
            <button class="btn btn-primary" onClick={() => Comm.get(fContext, AlgoGetMessage())}>
                AlgoGet
            </button>
        </div>
    );
};

export default Landing;