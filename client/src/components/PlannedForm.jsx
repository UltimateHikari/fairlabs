import { QueryMessage, SubmitMessage } from 'api/message/tasks';
import { AuthContext } from 'context';
import React, { useContext, useState, useEffect } from 'react';
import { Comm } from 'api/Comm';
import { PlannedIntent } from 'intents';
import { useNavigate } from 'react-router';


const PlannedForm = () => {
    const {fContext} = useContext(AuthContext)

    const [tasks, setTasks] = useState([])
    const [loading, setLoading] = useState(false);
    const navigate = useNavigate()

    const handleCheckChange = event => {
        tasks[event.target.value] = event.target.checked;
    };
    
    const submit = (event) => {
        event.preventDefault();
        console.log(tasks);
        let checked_tasks = tasks.filter(t => {return t.checked === true})
        console.log(checked_tasks);
        Comm.post(fContext, SubmitMessage(PlannedIntent, checked_tasks))
        .catch((err) => {
            console.log(err);
            })
            // TODO - check response?? not tonight bae
        navigate(-1)
    }

    useEffect(() => {
        setLoading(true);
        Comm.get(fContext, QueryMessage(PlannedIntent))
        .then(response => {
        console.log(response.data);
        setTasks(response.data.tasks.map(t => {return({n:t, checked:false})}));
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

    return(
        <div>
        <h3> Plannable</h3>
        <form onSubmit={submit}>
        {tasks.map(t => {
            return (
                <div class="form-check">
                <input class="form-check-input" type="checkbox" value={t.n} id={"submittask"+t.n} onClick={handleCheckChange}/>
                <label class="form-check-label" for="flexCheckDefault">
                    {t.n}
                </label>
                </div>
            )
        })}
        <button className="btn btn-primary" type="submit">Submit planned</button>
        </form>
        </div>
    )
};

export default PlannedForm