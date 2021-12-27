import React, { useContext, useEffect, useState } from 'react';
import CourseItem from "../components/CourseItem";
import {AuthContext} from "context";
import { MyCoursesMessage } from "api/message/tasks.js"
import { Comm } from 'api/Comm';


//получать курсы конкретного тела
const MyCourses = () => {
    const {fContext} = useContext(AuthContext);
    const [courses, setCourses] = useState([])
    const [loading, setLoading] = useState(false);
    
    useEffect(() => {
        setLoading(true);
        Comm.get(fContext, MyCoursesMessage())
        .then(response => {
        console.log(response.data.courses);
        setCourses(response.data.courses);
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
        <div className={'My Courses'}>
            <h1 style={{textAlign: 'center', color: '#008080FF'}}> My courses list</h1>
            {courses.map(c => <CourseItem key={c.id} course={{id: c.id, title: c.name, body: 'group: ' + c.cgroup}}/>)}
        </div>
    );
};

export default MyCourses;