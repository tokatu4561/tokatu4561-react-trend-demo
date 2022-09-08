import React from 'react'
import { MainLayout } from '@/components/Layout/MainLayout'

import FormikTestForm from '@/features/form-test/components/FormikTestForm'

const FormikTest = () => {
  return (
    <MainLayout title="react-hook-form test">
      <div className="flex justify-center items-center h-screen w-4/6">
        <FormikTestForm />
      </div>
    </MainLayout>
  )
}

export default FormikTest
