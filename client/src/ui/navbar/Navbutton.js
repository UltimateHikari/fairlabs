import React, {useContext, useState} from 'react';
import {Link, useNavigate} from "react-router-dom";
import MyButton from "../button/MyButton";
import {AuthContext} from "../../context";
import ModalNote from "../../components/ModalNote";
import Login from "../../pages/Login";
import { useCookies } from 'react-cookie';
import Navlist from './Navlist';

const Navbutton = () =>{
    const {isAuth, setIsAuth, setPerson} = useContext(AuthContext);

    const [isModalActive, setModalActive] = useState(false)
    const [cookies, setCookie, removeCookie] = useCookies(['faircookie']);


    const navigate = useNavigate()

    const logout = () => {
        setIsAuth(false)
        setPerson("")
        setModalActive(false)
        removeCookie('auth', {})
        removeCookie('context', {})
        navigate('/landing')
    }

    if(isAuth == true){
    return(
        <div class="d-flex">
            <button class="btn" onClick={logout} >Sign out</button>
        </div>
    )
    }
    return(
        <div class="d-flex">
            <button class="btn" onClick={() => setModalActive(true)} >Sign in</button>
            <ModalNote active={isModalActive} setActive={setModalActive}>
                    <Login/>
            </ModalNote>
        </div>
    )
};

export default Navbutton;