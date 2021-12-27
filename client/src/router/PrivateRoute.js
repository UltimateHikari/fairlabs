import React, {useContext} from 'react';
import {AuthContext} from "../context";
import {Navigate, Outlet, Route} from "react-router";

export default function PrivateRoute(props){

    const {isAuth, setIsAuth} = useContext(AuthContext);
    const {role, setRole} = useContext(AuthContext);

    function checkRole(){
        return role >= props.role;
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