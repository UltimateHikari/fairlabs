import React, {useContext} from 'react';
import {BrowserRouter, Navigate, Route, Routes} from "react-router-dom";
import About from "../pages/About";
import Profile from "../pages/Profile";
import {privateRoutes, publicRoutes} from "./index";
import Login from "../pages/Login";
import {AuthContext} from "../context";


const AppRouter = () => {

    const {isAuth, setIsAuth} = useContext(AuthContext);

    return (

            <Routes>
                <Route path="/about" element={<About/>}/>
                <Route path="/login" element={<Login/>}/>
                {/*{privateRoutes.map((route, i) =>*/}
                {/*    <Route key={i} component={route.component}*/}
                {/*           path={route.path}*/}
                {/*           exact={route.exact}/>*/}
                {/*)}*/}
                {/*{publicRoutes.map(route =>*/}
                {/*    <Route component={route.component}*/}
                {/*           path={route.path}*/}
                {/*           exact={route.exact}/>*/}
                {/*)}*/}
                <Route path="*" element={<Navigate to={'/landing'}/>}/>
            </Routes>
    );
};

export default AppRouter;