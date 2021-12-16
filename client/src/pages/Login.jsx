import React, {useState} from 'react';

const Login = () => {

    const[text, setText] = useState("Please log in")

    return (
        <div>
            <h1>{text}</h1>
            <button>Log in</button>
        </div>
    );
};

export default Login;