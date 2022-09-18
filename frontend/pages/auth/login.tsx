import React from 'react'

import { InputFiled } from '@/components/Form/InputFiled'
import { useForm } from 'react-hook-form'
import { Layout } from '@/features/auth/components/Layout'
import { Button } from '@/components/Button'
import { fetchJson } from '@/lib/api'
import { axios } from '@/lib/axios'

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
    const response = axios.post('/auth/login', {
      email: data.email,
      password: data.password,
    })
    console.log(response)
  }
  const onError = (errors, e) => console.log(errors, e)

  const onClickGetUser = () => {
    const response = axios.get('/auth/user')
    console.log(response)
  }

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
              type="text"
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
              type="password"
              placeholder="password"
              {...register('password', { required: true })}
            />
            {errors.password && (
              <span className="text-red-500">入力は必須です。</span>
            )}
          </div>
          <Button type="submit">送信</Button>
        </form>

        <div className="ml-4">
          <Button onClick={onClickGetUser}>ユーザー取得</Button>
        </div>
      </div>
    </Layout>
  )
}

export default LoginPage
