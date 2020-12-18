import React, {useEffect} from 'react';
import { useFormik } from 'formik';

function CreateEntity(props){

    const formik = useFormik({
        initialValues: {
            key: '',
            value: '',
            description: ''
        }, 
        onSubmit: values => {
            console.log(values)
        }
    })

    useEffect(() => {
        console.log(props.match.params.id)
    })

    return(
        <div>Hey</div>
    )
}

export default CreateEntity;