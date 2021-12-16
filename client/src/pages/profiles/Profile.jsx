import React, {useContext} from 'react';
import {AuthContext} from "../../context";
import MyButton from "../../ui/button/MyButton";
import StudentProfile from "./StudentProfile";
import TeacherProfile from "./TeacherProfile";
import AdminProfile from "./AdminProfile";

/* тут должна быть вся фигня про чела
*  + курсу
*  + уСпЕхИ(?)
*  + кнопочки */

const Profile = () => {

    const {role, setRole} = useContext(AuthContext);

    function getRoleProfile(){
        if (role === 1) {
            return <AdminProfile/>;
        } else if (role === 2) {
            return <StudentProfile/>;
        } else if (role === 3) {
            return <TeacherProfile/>;
        } else {

        }
    }

    return (
        getRoleProfile()
    );
};

export default Profile;