import React, { FC } from 'react'

interface Props {
  inputId: string
  type: string
  label: string
  placeholder: string
}

export const InputFiled: FC<Props> = ({
  inputId,
  type,
  label,
  placeholder,
}) => {
  return (
    <>
      <label
        className="block text-gray-700 text-sm font-bold mb-2"
        htmlFor={inputId}
      >
        {label}
      </label>
      <input
        id={inputId}
        className="shadow border rounded w-full py-2 px-3 text-gray-700"
        type={type}
        placeholder={placeholder}
      />
    </>
  )
}
