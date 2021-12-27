
import {BrowserRouter, Route, Routes, Redirect} from "react-router-dom";
import { Comm, FContext } from "./api/Comm";
import './styles/App.css';
import AppRouter from "./router/AppRouter";
import {useEffect, useState} from "react";
import {AuthContext} from "./context";
import Navbar from "./ui/navbar/Navbar";
import { useCookies } from 'react-cookie';

// roles
// 1 == student
// 2 == teacher
// 3 == admin

function App() {

    const [isAuth, setIsAuth] = useState(false);
    const [fContext, setFContext] = useState(FContext);
    const [cookies, setCookie] = useCookies(['faircookie']);

    const setPerson = (email) => {
        let newFContext = fContext
        newFContext.email = email
        console.log(email);
        setFContext(newFContext)
    }

    const setCourse = (course_id) => {
        let newFContext = fContext
        newFContext.course = course_id
        setFContext(newFContext)
    }

    useEffect( () => {
        if (typeof cookies.auth != undefined){
            setIsAuth(cookies.auth)
        }
        if (typeof cookies.context != undefined){
            setIsAuth(cookies.context)
        }
    }, [])

    return (
        <AuthContext.Provider value={{
            isAuth,
            setIsAuth,
            fContext,
            setPerson,
            setCourse
        }}>

            <BrowserRouter>
                <Navbar/>
                <AppRouter/>
            </BrowserRouter>

        </AuthContext.Provider>
    );
}

export default App;
