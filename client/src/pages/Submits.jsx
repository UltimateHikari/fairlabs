import PlannedForm from 'components/PlannedForm';
import SubmitForm from 'components/SubmitForm';
import React from 'react';
import { useNavigate } from 'react-router';
import Widget from "../ui/widget/Widget";

const Submits = () => {
    return(
        <div>
        <h1> Submits </h1>
        <div class="container">
            <div class="row">
                <div class="col-6"><SubmitForm/></div>
                <div class="col-6"><PlannedForm/></div>
            </div>
        </div>
        </div>
    )
};

export default Submits