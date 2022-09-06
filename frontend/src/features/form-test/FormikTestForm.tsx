import React from 'react'
import { MainLayout } from '@/components/Layout/MainLayout'
import { useFormik } from 'formik'
import * as yup from 'yup'

const initialValues = {
  aaa: '',
  bbb: '',
  ccc: '',
}

const formValidationSchema = yup.object().shape({
  aaa: yup.string().required(),
  bbb: yup.string().required(),
  ccc: yup.string().email(),
})

const FormikTestForm = () => {
  const handleFormSubmit = (data) => {}

  const { values, errors, touched, handleBlur, handleChange, handleSubmit } =
    useFormik({
      onSubmit: handleFormSubmit,
      initialValues,
      validationSchema: formValidationSchema,
    })

  console.log('renderd')

  return (
    <form onSubmit={handleSubmit}>
      Formik
      <div className="mb-4">
        <input
          name="aaa"
          className="p-2"
          placeholder="aaa"
          type="text"
          onBlur={handleBlur}
          onChange={handleChange}
          value={values.aaa}
          // error={!!touched.email && !!errors.email}
          // helperText={touched.email && errors.email}
        />
      </div>
      <div className="mb-4">
        <input
          className="p-2"
          type="text"
          placeholder="bbb"
          name="bbb"
          onBlur={handleBlur}
          onChange={handleChange}
          value={values.bbb}
        />
        {/* {errors.bbb && <p className="text-red-400">入力は必須です。</p>} */}
      </div>
      <div className="mb-4">
        <input
          name="ccc"
          className="p-2"
          placeholder="ccc"
          value={values.ccc}
          onBlur={handleBlur}
          onChange={handleChange}
        />
      </div>
      <button
        type="submit"
        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
      >
        送信
      </button>
    </form>
  )
}

export default FormikTestForm
