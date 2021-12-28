import React from 'react';

import Navlist from './Navlist';
import Navbutton from './Navbutton';


const Navbar = () => {

    return(
        <nav class="navbar navbar-expand-lg navbar-light bg-light">
        <div class="container-fluid">
            <a class="navbar-brand" href="/">Fairlabs</a>
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