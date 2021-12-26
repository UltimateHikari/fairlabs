import React from 'react';
import CourseItem from "../components/CourseItem";


//все курсы
//для админов кнопка редактирования
//для препода + кнопка фоллоу
const CoursesList = () => {
    return (
        <div className={'Courses'}>
            <h1 style={{textAlign: 'center'}}> my courses</h1>
            <CourseItem course={{id: 1, title: 'Course', body: 'i am coursedfxgcfhvgbhxcvbnmn,fcvgbhjmvgbhn'}}/>
            <CourseItem course={{id: 2, title: 'Course', body: 'i am course'}}/>
        </div>
    );
};

export default CoursesList;