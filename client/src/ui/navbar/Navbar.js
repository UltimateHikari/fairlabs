import React, {useContext, useState} from 'react';
import {Link, useNavigate} from "react-router-dom";
import MyButton from "../button/MyButton";
import {AuthContext} from "../../context";
import ModalNote from "../../components/ModalNote";
import Login from "../../pages/Login";
import { useCookies } from 'react-cookie';
import Navlist from './Navlist';
import Navbutton from './Navbutton';


const Navbar = () => {

    return(
        <nav class="navbar navbar-expand-lg navbar-light bg-light">
        <div class="container-fluid">
            <a class="navbar-brand" href="#">Navbar</a>
            <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse" id="navbarSupportedContent">
            <Navlist/>
            <Navbutton/>
            </div>
        </div>
        </nav>
    )
};

export default Navbar;