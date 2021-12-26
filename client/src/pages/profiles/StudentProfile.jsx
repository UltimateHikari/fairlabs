import React, {useContext, useState} from 'react';
import MyButton from "../../ui/button/MyButton";
import ModalNote from "../../components/ModalNote";
import Widget from "../../ui/widget/Widget";
import {AuthContext} from "../../context";
import AdminProfile from "./AdminProfile";


//отметить запланированное задание и сданное - отдельный страницы
//мне не нравится как оно выглядит
//убрать отсюда виджеты на страницу курса
const StudentProfile = () => {

    const {role, setRole} = useContext(AuthContext);

    const [isModalActive, setModalActive] = useState(false)

    function declarePriority(){
        console.log("declarePriority")
    }

    function getAdmin(){
        if (role !== 2){
            return <AdminProfile/>;
        }
    }

    return (
        <div className={"student_profile"}>
            <div className={"student_buttons"}>
                <button className={"profile_button"} onClick={declarePriority}>
                    DecPriority
                </button>

                <button className={"profile_button"} onClick={() => setModalActive(true)}>
                    modal
                </button>
                <ModalNote active={isModalActive} setActive={setModalActive}>
                    <p>Add passed tasks</p>

                </ModalNote>
            </div>
            <div className={"student_widgets"}>
                <Widget>
                    <p>My Goals</p>
                </Widget>
                <Widget>
                    <p>My Progress</p>
                </Widget>
            </div>

        </div>
    );
};

export default StudentProfile;