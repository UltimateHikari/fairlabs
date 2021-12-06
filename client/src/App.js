
import {BrowserRouter, Route, Routes, Redirect} from "react-router-dom";
import About from "./pages/About";
import axios from "axios";
import Comm from "./api/Comm";
import Navbar from "./ui/Navbar";
import './styles/App.css';
import AppRouter from "./router/AppRouter";
import {useState} from "react";
import {AuthContext} from "./context";

function App() {

    const [isAuth, setIsAuth] = useState(false);

    async function fetch(){
        const response = await Comm.getSth()
        console.log(response)
    }

    return (
        <AuthContext.Provider value={{
            isAuth,
            setIsAuth
        }}>

            <BrowserRouter>
                {/*<Navbar/>*/}
                <AppRouter/>
            </BrowserRouter>

            {/*<button onClick={fetch}>get</button>*/}
        </AuthContext.Provider>

    );
}

export default App;
