import React, {useEffect} from 'react';
import { useFormik } from 'formik';

function Main(props){

    const formik = useFormik({
        initialValues: {
            key: '',
            value: ''
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

export default Main;