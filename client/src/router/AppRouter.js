import React, {useContext} from 'react';
import {BrowserRouter, Navigate, Route, Routes} from "react-router-dom";
import About from "../pages/About";
import Profile from "../pages/profiles/Profile";
import Login from "../pages/Login";
import {AuthContext} from "../context";
import Landing from "../pages/Landing";
import Error from "../pages/Error";
import CoursesList from "../pages/CoursesList";
import CourseCreate from "../pages/CourseCreate";
import PrivateRoute from "./PrivateRoute";
import Security from "../pages/Security";
import MyCourses from "../pages/MyCourses";
import PublicRoute from "./PublicRoute";
import Course from "../pages/Course";
import CourseEdit from "../pages/CourseEdit";


const AppRouter = () => {

    return (
        <Routes>
            <Route path="/about" element={<About/>}/>
            <Route path="/landing" element={<Landing/>}/>

            <Route element={<PublicRoute/>}>
                <Route path="/login" element={<Login/>}/>
            </Route>

            <Route element={<PrivateRoute role={1}/>}>
                <Route path="/profile" element={<Profile/>}/>
            </Route>
            <Route element={<PrivateRoute role={1}/>}>
                <Route exact path="/courses" element={<MyCourses/>}/>
            </Route>
            <Route element={<PrivateRoute role={1}/>}>
                <Route exact path="/courses/:id" element={<Course/>}/>
            </Route>
            <Route element={<PrivateRoute role={2}/>}>
                <Route path="/create_course" element={<CourseCreate/>}/>
            </Route>
            <Route element={<PrivateRoute role={2}/>}>
                <Route path="/edit_course" element={<CourseEdit/>}/>
            </Route>
            <Route element={<PrivateRoute role={2}/>}>
                <Route path="/courses_list" element={<CoursesList/>}/>
            </Route>



            <Route path="/security" element={<Security/>}/>
            <Route path="/error" element={<Error/>}/>
            <Route path="/" element={<Navigate to={'/landing'}/>}/>
            <Route path="*" element={<Navigate to={'/error'}/>}/>
        </Routes>
    );
};

export default AppRouter;