import React from 'react'
import { MainLayout } from '@/components/Layout/MainLayout'
import { InputFiled } from '@/components/Form/InputFiled'
import { useForm } from 'react-hook-form'

type Inputs = {
  aaa: string
  bbb: string
}

const FormTest = () => {
  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm<Inputs>()

  const onSubmit = (data) => {
    console.log(data)
  }

  return (
    <MainLayout title="react-hook-form test">
      react-hook-form test
      <div className="flex justify-center items-center h-screen w-4/6">
        <form onSubmit={handleSubmit(onSubmit)}>
          <div className="mb-4">
            <InputFiled
              inputId="aaa"
              label="AAA"
              type="text"
              placeholder="aaa"
              {...register('aaa', { required: true })}
            />
            {errors.aaa && (
              <span className="text-red-500">This field is required</span>
            )}
          </div>
          <div className="mb-4">
            <InputFiled
              inputId="bbb"
              label="BBB"
              type="text"
              placeholder="bbbb"
              {...register('bbb')}
            />
          </div>
          <div className="mb-4">
            <InputFiled
              inputId="ccc"
              label="CCC"
              type="text"
              placeholder="ccc"
            />
          </div>
          <div className="mb-4">
            <InputFiled
              inputId="ddd"
              label="DDD"
              type="text"
              placeholder="ddd"
            />
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

export default FormTest
