import React, {useContext, useState} from 'react';
import {AuthContext} from "../context";
import { useCookies } from 'react-cookie';



//поле для ввода почты
//студ ент/учитель?
const Login = () => {
    const [email, setEmail] = useState('')
    
    const handleEmailChange = event => {
        setEmail(event.target.value)
    };

    const {isAuth, fContext, setIsAuth, setPerson} = useContext(AuthContext);
    const [cookies, setCookie] = useCookies(['faircookie']);

    const login = (event) => {
        event.preventDefault();
        setIsAuth(true);
        setPerson(email);

        setCookie('auth', isAuth,{path: '/'})
        setCookie('context', fContext,{path: '/'})
    }

    return (
        <div className={'login'}>
            <form onSubmit={login}>
                <h1 className={'h1'}>Please log in</h1>
                <div>
                    <label>Email address</label>
                    <input
                    type="email"
                    name="email"
                    placeholder="Enter email"
                    onChange={handleEmailChange}
                    value={email}
                    />
                </div>
                <button className="btn" type="submit">Log in</button>
            </form>
        </div>
    );
};

export default Login;