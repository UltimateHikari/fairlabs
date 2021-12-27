import React, {useContext} from 'react';
import {Link} from "react-router-dom";

import {AuthContext} from "../../context";
import { StudentRole, TeacherRole } from 'roles';

const Navlist = () =>{
    const {isAuth, fContext} = useContext(AuthContext);

    if(isAuth === true && fContext.role === TeacherRole){
        return(
            <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                <li class="nav-item">
                <Link className="nav-link active" to="/about">About</Link>
                </li>
                <li class="nav-item">
                <Link className="nav-link active" to="/landing">Home</Link>
                </li>
                <li class="nav-item">
                <Link className="nav-link active" to="/courses">Courses</Link>
                </li>
                <li class="nav-item">
                <Link className="nav-link active" to="/profile">Admin Page</Link>
                </li>
                <li class="nav-item">
                <Link className="nav-link active" to="/create_course">New Course</Link>
                </li>
            </ul>
        )
    }

    if(isAuth === true && fContext.role === StudentRole){
        return(
            <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                <li class="nav-item">
                <Link className="nav-link active" to="/about">About</Link>
                </li>
                <li class="nav-item">
                <Link className="nav-link active" to="/landing">Home</Link>
                </li>
                <li class="nav-item">
                <Link className="nav-link active" to="/courses">Courses</Link>
                </li>
            </ul>
        )
    }

    return(
        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
            <li class="nav-item">
            <Link className="nav-link active" to="/about">About</Link>
            </li>
            <li class="nav-item">
            <Link className="nav-link active" to="/landing">Home</Link>
            </li>
        </ul>
    )
    
};

export default Navlist;