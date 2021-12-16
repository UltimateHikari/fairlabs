import About from "../pages/About";
import Profile from "../pages/profiles/Profile";
import Landing from "../pages/Landing";
import Login from "../pages/Login";
import {Navigate} from "react-router-dom";
import React from "react";

// export const routes = [
//     {path: '/about', component: About, exact: true},
//     {path: '/profile', component: Profile, exact: true},
//     {path: '/landing', component: Landing, exact: true},
// ]

export const privateRoutes = {
    "About":{path: '/about', component: About, exact: true},
    "Profile":{path: '/profile', component: Profile, exact: true},
}

export const publicRoutes = [
    {path: '/login', component: Login, exact: true},
    {path: '/landing', component: Landing, exact: true},
    {path: '*', component: <Navigate to={'/landing'}/>, exact: true},
]