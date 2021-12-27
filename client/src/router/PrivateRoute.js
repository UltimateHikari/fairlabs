import React, {useContext} from 'react';
import {AuthContext} from "../context";
import {Navigate, Outlet, Route} from "react-router";

export default function PrivateRoute(props){

    const {isAuth} = useContext(AuthContext);
    const {fContext} = useContext(AuthContext);

    function checkRole(){
        return fContext.role >= props.role;
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