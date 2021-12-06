import React from 'react';
import {Link} from "react-router-dom";

const Navbar = () => {
    return (
        <div className='navbar'>
            {/*<button onClick={logout}>*/}
            {/*    Выйти*/}
            {/*</button>*/}
            <div className="navbar_links">
                <Link to="/about">About</Link>
                <Link to="/profile">Profile</Link>
            </div>
        </div>
    );
};

export default Navbar;