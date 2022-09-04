import React, { FC } from 'react'

interface Props {
  inputId: string
  type: string
  label: string
  placeholder: string
}

export const InputFiled: FC<Props> = React.forwardRef(
  ({ inputId, type, label, placeholder, ...props }, ref) => {
    return (
      <>
        <label
          className="block text-gray-700 text-sm font-bold mb-2"
          htmlFor={inputId}
        >
          {label}
        </label>
        <input
          className="shadow border rounded w-full py-2 px-3 text-gray-700"
          type={type}
          placeholder={placeholder}
          ref={ref as React.ForwardedRef<HTMLInputElement>}
          {...props}
        />
      </>
    )
  }
)
