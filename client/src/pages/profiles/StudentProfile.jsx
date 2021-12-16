import React from 'react';
import MyButton from "../../ui/button/MyButton";

const StudentProfile = () => {

    function declarePriority(){
        console.log("declarePriority")
    }

    return (
        <div>
            <div>
                <MyButton onClick={declarePriority}>
                    DecPriority
                </MyButton>
            </div>
        </div>
    );
};

export default StudentProfile;