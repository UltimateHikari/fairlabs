import React from 'react';
import MyButton from "../ui/button/MyButton";

//add tasks
//add group
//add teacher(?)

const CreateCourse = () => {

    function setTasks(){
        console.log("setTasks")
    }

    function setGroup(){
        console.log("setGroup")
    }
//или как это вообще
    return (
        <div>
            <MyButton onClick={setTasks}>
                Set tasks
            </MyButton>

            <MyButton onClick={setGroup}>
                Set Group
            </MyButton>
        </div>
    );
};

export default CreateCourse;