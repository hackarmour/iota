import React from 'react';
import {useFormik} from 'formik';

function CreateEntity() {
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
        <form onSubmit={formik.handleSubmit}>
            <label htmlFor="name">Name</label>
            <input
                id="name"
                name="name"
                type="text"
                onChange={formik.handleChange}
                value={formik.values.name}
            />

            <label htmlFor="description">Description</label>
            <input
                id="description"
                name="description"
                type="text"
                onChange={formik.handleChange}
                value={formik.values.description}
            />
        <button type="submit">Submit</button>
       </form>
    )
}

export default CreateEntity;