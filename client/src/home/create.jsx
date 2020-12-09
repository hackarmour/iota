import React from 'react';
import {useFormik} from 'formik';
import { Form } from '../components/create';

function CreateProject() {
    const formik = useFormik({
        initialValues: {
            name: '',
            description: ''
        },
        onSubmit: values => {
            console.log(values)
        }
    })
    return (
        <Form onSubmit={formik.handleSubmit}>
            <input
                id="name"
                name="name"
                type="text"
                placeholder="name"
                onChange={formik.handleChange}
                value={formik.values.name}
            />
            <br />
            <input
                id="description"
                name="description"
                placeholder="description"
                type="text"
                onChange={formik.handleChange}
                value={formik.values.description}
            />
            <br />
        <button type="submit">Submit</button>
       </Form>
    )
}

export default CreateProject;