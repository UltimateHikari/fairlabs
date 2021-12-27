import { QueueMessage } from 'api/message/tasks';
import { AuthContext } from 'context';
import React, { useContext, useState, useEffect } from 'react';
import { Comm } from 'api/Comm';

const Queue = () => {
    const {fContext} = useContext(AuthContext)

    const [queue, setQueue] = useState([])
    const [loading, setLoading] = useState(false);


    useEffect(() => {
        setLoading(true);
        Comm.get(fContext, QueueMessage())
        .then(response => {
        console.log(response.data);
        setQueue(response.data.queue);
        })
        .catch((err) => {
        console.log(err);
        })
        .finally(() => {
        setLoading(false);
        });
    },[])

    if (loading) {
        return <p>Data is loading...</p>;
      }

    return (
        <div>
        <h1>Queue</h1>
        {queue.map(e => {
            if(e.priority === true){
                return <p><strong>{e.name}</strong></p>
            }else{
                return <p>{e.name}</p>
            }
        })
        }
        </div>
    );
};

export default Queue;