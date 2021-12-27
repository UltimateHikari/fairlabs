import React from 'react';
import {useNavigate} from "react-router-dom";
import MyButton from "../../ui/button/MyButton";

//линк на все курсы с возможностью редактирования
const AdminProfile = () => {

    const navigate = useNavigate()

    return (
        <div>
            <MyButton onClick={() => navigate('/courses_list')}>
                admin
            </MyButton>
        </div>
    );
};

export default AdminProfile;