import React from 'react';
import Widget from "../ui/widget/Widget";


//страница инфы по курсу
//виджеты со своими успехами (Если студент)
//очередь
//кнопка для общих статов (CourseStats)
const Course = (props) => {
    return (
        <div>
            <div className={'student_widgets'}>
                <Widget>
                    <p>My Goals</p>
                </Widget>
                <Widget>
                    <p>My Progress</p>
                </Widget>
            </div>
            <div className={'course_goals'}>

            </div>
            <div className={'queue'}>

            </div>
            <div className={'course_data'}>

            </div>
        </div>
    );
};

export default Course;