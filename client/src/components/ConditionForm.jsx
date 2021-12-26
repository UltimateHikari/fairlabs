import React, {useState} from 'react';
import MyInput from "../ui/input/MyInput";
import MyButton from "../ui/button/MyButton";

//страница условия для курсов с подставлением example_data из запроса
//кнопка снизу добавить еще ряд + удалить
const ConditionForm = () => {

    const [formValues, setFormValues] = useState([{name: '', email: ''}])

    let handleChanges = (i, e) => {
        let newFormValues = [...formValues];
        newFormValues[i][e.target.name] = e.target.value;
        setFormValues(newFormValues)
    }

    let addFormFields = () => {
        setFormValues([...formValues, {name: '', email: ''}])
    }

    let removeFormFields = (i) => {
        let newFormValues = [...formValues];
        newFormValues.splice(i, 1);
        setFormValues(newFormValues)
    }

    const handleSubmit = (event) => {
        event.preventDefault();
        //todo
    }

    return (
        <div className={''}>
            {formValues.map((element, index) => (
                <div className={'form_inline'} key={index}>
                    <input type={'name'} placeholder={'enter name'} onChange={e => handleChanges(index, e)}/>
                    <input type={'email'} placeholder={'enter email'} onChange={e => handleChanges(index, e)}/>
                    {index ?
                        <button className={'btn cond'} onClick={() => removeFormFields()}>delete</button>
                        : null
                    }
                </div>
            ))}
            <div className={'form_buttons'}>
                <button className={'btn cond'} onClick={() => addFormFields()}>Add</button>
            </div>
        </div>
    );
};

export default ConditionForm;