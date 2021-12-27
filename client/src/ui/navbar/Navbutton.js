import React, {useContext, useState} from 'react';
import {useNavigate} from "react-router-dom";
import {AuthContext} from "../../context";
import ModalNote from "../../components/ModalNote";
import Login from "../../pages/Login";
import { useCookies } from 'react-cookie';


const Navbutton = () =>{
    const {isAuth, setIsAuth, setPerson, fContext} = useContext(AuthContext);

    const [isModalActive, setModalActive] = useState(false)
    const {removeCookie} = useCookies(['faircookie']);


    const navigate = useNavigate()

    const logout = () => {
        setIsAuth(false)
        setPerson("")
        setModalActive(false)
        removeCookie('auth', {})
        removeCookie('context', {})
        navigate('/landing')
    }

    if(isAuth === true){
    return(
        <div class="d-flex">
            <span class="navbar-text">
                Logged in as {fContext.email}
            </span>
            <button class="btn me-2" onClick={logout} >Sign out</button>
        </div>
    )
    }
    return(
        <div class="d-flex">
            <button class="btn me-2" onClick={() => setModalActive(true)} >Sign in</button>
            <ModalNote active={isModalActive} setActive={setModalActive}>
                    <Login/>
            </ModalNote>
        </div>
    )
};

export default Navbutton;