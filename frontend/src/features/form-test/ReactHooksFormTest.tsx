import React from 'react'
import { useForm } from 'react-hook-form'

type Inputs = {
  aaa: string
  bbb: string
  ccc: string
}

const ReactHookFormTestForm = () => {
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

  console.log('renderd')

  return (
    <form onSubmit={handleSubmit(onSubmit, onError)}>
      <div className="mb-4">
        <input
          className="p-2"
          type="text"
          placeholder="aaa"
          {...register('aaa')}
        />
      </div>
      <div className="mb-4">
        <input
          className="p-2"
          type="text"
          placeholder="bbb"
          {...register('bbb', { required: true })}
        />
        {errors.bbb && <p className="text-red-400">入力は必須です。</p>}
      </div>
      <div className="mb-4">
        <input
          id="name"
          className="p-2"
          placeholder="ccc"
          {...register('ccc', { required: true, maxLength: 5 })}
        />
        {errors.ccc && errors.ccc.type === 'required' && (
          <p className="text-red-400">入力は必須です</p>
        )}
        {errors.ccc && errors.ccc.type === 'maxLength' && (
          <p className="text-red-400">5文字以内で入力してください。</p>
        )}
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

export default ReactHookFormTestForm
