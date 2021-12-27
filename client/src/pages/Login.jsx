import React, {useContext, useState} from 'react';
import {AuthContext} from "../context";
import { useCookies } from 'react-cookie';
import { useNavigate } from 'react-router-dom'
import { StudentRole, TeacherRole } from 'roles';


//поле для ввода почты
//студ ент/учитель?
const Login = () => {
    const [email, setEmail] = useState('')
    const [checked, setChecked] = useState('')
    const navigate = useNavigate();
    
    const handleEmailChange = event => {
        setEmail(event.target.value)
    };
    const handleCheckChange = event => {
        setChecked(event.target.value)
    };

    const {isAuth, fContext, setIsAuth, setPerson} = useContext(AuthContext);
    const [cookies, setCookie] = useCookies(['faircookie']);

    const login = (event) => {
        event.preventDefault();
        setIsAuth(true);
        if(checked){
            setPerson(email, TeacherRole);
        }else{
            setPerson(email, StudentRole);
        }

        setCookie('auth', isAuth,{path: '/'})
        setCookie('context', fContext,{path: '/'})
        navigate('/courses')
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
                <div class="form-check form-switch">
                    <input 
                        class="form-check-input"
                        type="checkbox"
                        id="teacherSwitch" 
                        onChange={handleCheckChange}/>
                    <label
                        class="form-check-label"
                        for="teacherSwitch">
                            Я преподаватель (trust me lol)
                    </label>
                </div>
                <button className="btn" type="submit">Log in</button>
            </form>
        </div>
    );
};

export default Login;