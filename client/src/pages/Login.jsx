import React, {useContext, useState} from 'react';
import MyButton from "../ui/button/MyButton";
import MyInput from "../ui/input/MyInput";
import {AuthContext} from "../context";


//поле для ввода почты
//студ ент/учитель?
const Login = () => {
    const [email, setEmail] = useState('')
    
    const handleEmailChange = event => {
        setEmail(event.target.value)
    };

    const {isAuth, setIsAuth, setPerson} = useContext(AuthContext);

    const login = (event) => {
        event.preventDefault();
        setIsAuth(true);
        setPerson(email);
        localStorage.setItem('auth', 'true')
    }

    console.log(isAuth);

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