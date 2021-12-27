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
        if (role === 3) {
            return <div><StudentProfile/> <AdminProfile/></div>;
        } else if (role === 2) {
            return <div><TeacherProfile/> <AdminProfile/></div>;
        } else if (role === 1) {
            return <div><StudentProfile/></div>;
        } else {

        }
    }

    return (
        getRoleProfile()
    );
};

export default Profile;