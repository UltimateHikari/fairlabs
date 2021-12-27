import PlannedForm from 'components/PlannedForm';
import SubmitForm from 'components/SubmitForm';
import React from 'react';
import { useNavigate } from 'react-router';
import Widget from "../ui/widget/Widget";

const Submits = () => {
    return(
        <div>
        <h1> Submits </h1>
        <SubmitForm/>
        <PlannedForm/>
        </div>
    )
};

export default Submits