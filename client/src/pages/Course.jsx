import React from 'react';
import { useNavigate } from 'react-router';
import Widget from "../ui/widget/Widget";


//страница инфы по курсу
//виджеты со своими успехами (Если студент)
//очередь
//кнопка для общих статов (CourseStats)
const Course = (props) => {
    const navigate = useNavigate()

    const toQueue = () => {navigate('/queue')}
    const toSubmits = () => {navigate('/submits')}

    return (
        <div>
            <div className={'student_widgets'}>
                <Widget>
                    <p>My Goals {props.id}</p>
                </Widget>
                <Widget>
                    <p>My Progress</p>
                </Widget>
            </div>
            <button class="btn" onClick={toQueue}> To Queue </button>
            <button class="btn" onClick={toSubmits}> To Submits</button>
            <div className={'course_data'}>

            </div>
        </div>
    );
};

export default Course;