import React, {useContext, useState} from 'react';
import MyButton from "../ui/button/MyButton";
import MyInput from "../ui/input/MyInput";
import {useNavigate} from "react-router-dom";
import ConditionForm from "../components/ConditionForm";

//add tasks
//add group
//add teacher(?)


//переход на CourseEdit после указания названия (группы/препода?)
//ConditionForm
//форма выбора алгоритма
//кнопка подтвердить
const CourseCreate = ({props}) => {

    const navigate = useNavigate()

    const handleSubmit = (event) => {
        event.preventDefault();
        //todo
    }

//или как это вообще
    return (
        <div>
            <h1 className={'h1'}>Create Course</h1>
            <form onSubmit={handleSubmit}>
                <div className={'create_form'}>
                    <p>Default information</p>
                    <MyInput type={'name'} placeholder={'enter name'}/>
                    <MyInput type={'group'} placeholder={'enter group'}/>
                    <p>Conditions</p>
                    <ConditionForm/>
                    <button className={'submit_btn'}>Submit</button >
                </div>
            </form>
        </div>
    );
};

export default CourseCreate;