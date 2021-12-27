import React, {useContext, useState} from 'react';
import {Link, useNavigate} from "react-router-dom";
import MyButton from "../button/MyButton";
import {AuthContext} from "../../context";
import ModalNote from "../../components/ModalNote";
import Login from "../../pages/Login";

const Navbar = () => {

    const {isAuth, setIsAuth} = useContext(AuthContext);

    const [isModalActive, setModalActive] = useState(false)

    const navigate = useNavigate()

    const logout = () => {
        setIsAuth(false);
        setModalActive(false)
        localStorage.removeItem('auth')
        navigate('/landing')
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
                    <Link className="li a" to="/create_course">CCreate</Link>
                </div>
            </div>
            :
            <div className='navbar'>
                <MyButton onClick={() => setModalActive(true)}>
                    Log in
                </MyButton>
                <div className="li">
                    <Link className="li a" to="/about">About</Link>
                    <Link className="li a" to="/landing">Home</Link>
                </div>
                <ModalNote active={isModalActive} setActive={setModalActive}>
                    <Login/>
                </ModalNote>
            </div>
    );
};

export default Navbar;