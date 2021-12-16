import React, {useContext, useState} from 'react';
import {AuthContext} from "../context";
import ModalNote from "../components/ModalNote";
import {Navigate, Outlet, Route} from "react-router";
import About from "../pages/About";
import Profile from "../pages/profiles/Profile";
import Login from "../pages/Login";
import Landing from "../pages/Landing";
import Error from "../pages/Error";
import Courses from "../pages/Courses";
import CreateCourse from "../pages/CreateCourse";

export default function PrivateRoute(props){

    const {isAuth, setIsAuth} = useContext(AuthContext);
    const {role, setRole} = useContext(AuthContext);

    function checkRole(){
        return role <= props.role;
    }

    return isAuth ? (
        checkRole() ?
            <Outlet/>
            :
            <Navigate to="/security" replace />
    ) : (
        <Navigate to="/login" replace />
    );
}