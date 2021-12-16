import React, {useContext, useState} from 'react';
import {Link} from "react-router-dom";
import MyButton from "../button/MyButton";
import {AuthContext} from "../../context";

const Navbar = () => {

    const {isAuth, setIsAuth} = useContext(AuthContext);

    const login = () => {
        setIsAuth(true);
    }

    const logout = () => {
        setIsAuth(false);
    }

    return (
        isAuth ?
            <div className='navbar'>
                <MyButton onClick={logout}>
                    Log out
                </MyButton>
                <div className="li">
                    {/*мб сделать faq вместо about*/}
                    <Link className="li a" to="/about">About</Link>
                    <Link className="li a" to="/courses">Courses</Link>
                    <Link className="li a" to="/profile">Profile</Link>
                    <Link className="li a" to="/landing">Home</Link>
                    <Link className="li a" to="/c_create">CCreate</Link>
                </div>
            </div>
            :
            <div className='navbar'>
                <MyButton onClick={login}>
                    Log in
                </MyButton>
                <div className="li">
                    <Link className="li a" to="/about">About</Link>
                    <Link className="li a" to="/landing">Home</Link>
                </div>
            </div>
    );
};

export default Navbar;