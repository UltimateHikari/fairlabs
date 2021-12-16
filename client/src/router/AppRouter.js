import React, {useContext} from 'react';
import {BrowserRouter, Navigate, Route, Routes} from "react-router-dom";
import About from "../pages/About";
import Profile from "../pages/profiles/Profile";
import {privateRoutes, publicRoutes} from "./index";
import Login from "../pages/Login";
import {AuthContext} from "../context";
import Landing from "../pages/Landing";
import Error from "../pages/Error";
import Courses from "../pages/Courses";
import CreateCourse from "../pages/CreateCourse";
import PrivateRoute from "./PrivateRoute";
import Security from "../pages/Security";


const AppRouter = () => {

    const {isAuth, setIsAuth} = useContext(AuthContext);

    console.log(isAuth)

    return (
        <Routes>
            <Route path="/about" element={<About/>}/>
            <Route path="/login" element={<Login/>}/>
            <Route path="/landing" element={<Landing/>}/>

            <Route element={<PrivateRoute role={3}/>}>
                <Route path="/profile" element={<Profile/>}/>
            </Route>

            <Route element={<PrivateRoute role={1}/>}>
                <Route path="/c_create" element={<CreateCourse/>}/>
            </Route>

            <Route path="/courses" element={<PrivateRoute component={<Courses/>}/>}/>

            <Route path="/security" element={<Security/>}/>
            <Route path="/error" element={<Error/>}/>
            <Route path="/" element={<Navigate to={'/landing'}/>}/>
            <Route path="*" element={<Navigate to={'/error'}/>}/>
        </Routes>
    );
};

export default AppRouter;