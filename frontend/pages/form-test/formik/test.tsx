import React from 'react'
import { MainLayout } from '@/components/Layout/MainLayout'
import { useFormik } from 'formik'

const initialValues = {
  aaa: '',
  bbb: '',
  ccc: '',
}

const formSchema = {}

const FormikTest = () => {
  const handleFormSubmit = (data) => {}

  const { values, errors, touched, handleBlur, handleChange, handleSubmit } =
    useFormik({
      onSubmit: handleFormSubmit,
      initialValues,
      validationSchema: formSchema,
    })

  return (
    <MainLayout title="react-hook-form test">
      <div className="flex justify-center items-center h-screen w-4/6">
        <form className="content" onSubmit={handleSubmit}>
          Formik
          <div className="mb-4">
            <input
              name="email"
              placeholder="exmple@mail.com"
              type="email"
              onBlur={handleBlur}
              onChange={handleChange}
              value={values.email}
              // error={!!touched.email && !!errors.email}
              // helperText={touched.email && errors.email}
            />
          </div>
          <div className="mb-4">
            <input className="p-2" type="text" placeholder="bbb" />
            {errors.bbb && <p className="text-red-400">入力は必須です。</p>}
          </div>
          <div className="mb-4">
            <input id="name" className="p-2" placeholder="ccc" />
          </div>
          <button
            type="submit"
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
          >
            送信
          </button>
        </form>
      </div>
    </MainLayout>
  )
}

export default FormikTest
