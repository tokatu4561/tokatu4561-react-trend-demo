import React from 'react'

import { InputFiled } from '@/components/Form/InputFiled'
import { useForm } from 'react-hook-form'
import { Layout } from '@/features/auth/components/Layout'
import { Button } from '@/components/Button'

type Inputs = {
  email: string
  password: string
}

const LoginPage = () => {
  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm<Inputs>()

  const onSubmit = (data) => {
    console.log(data)
  }
  const onError = (errors, e) => console.log(errors, e)

  return (
    <Layout title="login">
      <div className="flex justify-center items-center h-screen">
        <form
          className="w-full max-w-md"
          onSubmit={handleSubmit(onSubmit, onError)}
        >
          <div className="mb-4">
            <InputFiled
              inputId="email"
              label="email"
              type="password"
              placeholder="email"
              {...register('email', { required: true })}
            />
            {errors.email && (
              <span className="text-red-500">入力は必須です。</span>
            )}
          </div>
          <div className="mb-4">
            <InputFiled
              inputId="password"
              label="password"
              type="text"
              placeholder="password"
              {...register('password', { required: true })}
            />
            {errors.password && (
              <span className="text-red-500">入力は必須です。</span>
            )}
          </div>
          <Button type="submit">送信</Button>
        </form>
      </div>
    </Layout>
  )
}

export default LoginPage
