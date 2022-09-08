import React from 'react'
import { MainLayout } from '@/components/Layout/MainLayout'
import ReactHookFormTestForm from '@/features/form-test/components/ReactHooksFormTest'

const FormTest = () => {
  return (
    <MainLayout title="react-hook-form test">
      <div className="flex justify-center items-center h-screen w-4/6">
        <ReactHookFormTestForm />
      </div>
    </MainLayout>
  )
}

export default FormTest
