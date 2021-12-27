import React, {useContext} from 'react';
import {AuthContext} from "../../context";
import MyButton from "../../ui/button/MyButton";
import StudentProfile from "./StudentProfile";
import TeacherProfile from "./TeacherProfile";
import AdminProfile from "./AdminProfile";
import { StudentRole, TeacherRole } from 'roles';

/* тут должна быть вся фигня про чела
*  + курсу
*  + уСпЕхИ(?)
*  + кнопочки */

const Profile = () => {

    const {fContext} = useContext(AuthContext);

    function getRoleProfile(){
        if (fContext.role === TeacherRole) {
            return <div><TeacherProfile/> <AdminProfile/></div>;
        } else if (fContext.role === StudentRole) {
            return <div><StudentProfile/></div>;
        } else {

        }
    }

    return (
        getRoleProfile()
    );
};

export default Profile;