import React, {useContext, useState} from 'react';
import MyButton from "../ui/button/MyButton";
import MyInput from "../ui/input/MyInput";
import {AuthContext} from "../context";


//поле для ввода почты
//студ ент/учитель?
const Login = () => {

    const {isAuth, setIsAuth} = useContext(AuthContext);

    const login = (event) => {
        event.preventDefault();
        setIsAuth(true);
        localStorage.setItem('auth', 'true')
    }

    console.log(isAuth);

    return (
        <div className={'login'}>
            <form onSubmit={login}>
                <h1 className={'h1'}>Please log in</h1>
                <MyInput type="e-mail" placeholder="Enter e-mail"/>
                <button className={'btn'}>Log in</button>
            </form>
        </div>
    );
};

export default Login;