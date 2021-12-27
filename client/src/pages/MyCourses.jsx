import React from 'react';
import Course from "./Course";
import CourseItem from "../components/CourseItem";


//получать курсы конкретного тела
const MyCourses = () => {
    return (
        <div className={'Courses'}>
            <h1 style={{textAlign: 'center', color: '#008080FF'}}> My courses</h1>
            <CourseItem course={{id: 1, title: 'Course', body: 'i am coursedfxgcfhvgbhxcvbnmn,fcvgbhjmvgbhn'}}/>
            <CourseItem course={{id: 2, title: 'Course', body: 'i am course'}}/>
        </div>
    );
};

export default MyCourses;