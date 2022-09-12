import { todoListAtom } from '@/stores/atoms/jotai/todoListAtim'
import { atom, useAtom } from 'jotai'
import React, { useState } from 'react'
import { Todo } from '../types'

const addTodoAtom = atom(null, (get, set, title) => {
  const todoList = get(todoListAtom)
  set(todoListAtom, [...todoList, { title: title } as Todo])
})

export const JotaiCreateForm = () => {
  const [title, setTitle] = useState('')

  const [, addTodo] = useAtom(addTodoAtom)

  const onChange = (event) => {
    setTitle(event.target.value)
  }

  const onClick = () => {
    addTodo(title)
  }

  return (
    <>
      <label>
        タスク名:
        <input
          type="text"
          value={title}
          onChange={onChange}
          name="title"
          style={{ margin: 12 }}
        />
      </label>
      <button onClick={onClick}>追加</button>
    </>
  )
}
