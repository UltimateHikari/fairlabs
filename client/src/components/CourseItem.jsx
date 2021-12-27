import React, {useContext} from 'react';
import {useNavigate} from "react-router-dom";
import {AuthContext} from "../context";
import MyButton from "../ui/button/MyButton";

const CourseItem = (props) => {

    const navigate = useNavigate()
    const {role, setRole} = useContext(AuthContext);

    function getButtons({}) {
        let id = props.course.id
        if (role >= 2) {
            return (
                <div>
                    <MyButton onClick={() => navigate('/courses/' + id.toString())}>
                        check
                    </MyButton>
                    <MyButton>
                        edit
                    </MyButton>
                </div>
            );
        } else {
            return (
                <div>
                    <MyButton onClick={() => navigate('/courses/' + id.toString())}>
                        check
                    </MyButton>
                </div>
            );
        }
    }

    return (
        <div className={"course"}>
            <div className={"course_content"}>
                <strong>{props.course.id} {props.course.title}</strong>
                <div>
                    {props.course.body}
                </div>
            </div>
            <div className={"btns"}>
                {getButtons(props.course.id)}
            </div>
        </div>
    );
};

export default CourseItem;