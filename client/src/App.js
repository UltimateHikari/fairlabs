
import {BrowserRouter, Route, Routes, Redirect} from "react-router-dom";
import About from "./pages/About";
import axios from "axios";
import { Comm } from "./api/Comm";
import './styles/App.css';
import AppRouter from "./router/AppRouter";
import {useEffect, useState} from "react";
import {AuthContext} from "./context";
import Navbar from "./ui/navbar/Navbar";

// roles
// 1 == student
// 2 == teacher
// 3 == admin

function App() {

    const [isAuth, setIsAuth] = useState(false);
    const [role, setRole] = useState(2);
    const [group, setGroup] = useState(19201);

    useEffect( () => {
        if (localStorage.getItem('auth')){
            setIsAuth(true)
        }
    }, [])

    async function fetch(){
        const response = await Comm.getSth()
        console.log(response)
    }
//добавить в конеткст емейл сетЕмейл?
    return (

        <AuthContext.Provider value={{
            isAuth,
            setIsAuth,
            role,
            setRole,
            group,
            setGroup
        }}>

            <BrowserRouter>
                <Navbar/>
                <AppRouter/>
            </BrowserRouter>

        </AuthContext.Provider>
        // <div>
        //     <button onClick={fetch}>get</button>
        // </div>
    );
}

export default App;
