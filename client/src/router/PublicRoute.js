import React, {useContext} from 'react';
import {AuthContext} from "../context";
import {Navigate, Outlet} from "react-router";

const PublicRoute = () => {

    const {isAuth, setIsAuth} = useContext(AuthContext);

    return isAuth ? (
        <Navigate to="/profile" replace />
    ) : (
        <Outlet/>
    );
};

export default PublicRoute;