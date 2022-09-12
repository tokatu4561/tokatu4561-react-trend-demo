import { MainLayout } from '@/components/Layout/MainLayout'
import CreateForm from '@/features/todo/components/CreateForm'
import SearchForm from '@/features/todo/components/SearchForm'
import { TodoList } from '@/features/todo/components/TodoList'
import { Todo } from '@/features/todo/types'

import { useAtom, atom } from 'jotai'
import React, { useState } from 'react'

const todoListAtom = atom<Todo[]>([{ title: 'aaa' }])

const addTodoAtom = atom(null, (get, set, title) => {
  const todoList = get(todoListAtom)
  set(todoListAtom, [...todoList, { title: title } as Todo])
})

const JotaiTest = () => {
  const [todoList, setTodoList] = useAtom(todoListAtom)

  const [, addTodo] = useAtom(addTodoAtom)

  const [title, setTitle] = useState('')

  const onChange = (event) => {
    setTitle(event.target.value)
  }

  const onClick = () => {
    addTodo(title)
  }

  return (
    <MainLayout title="stripe">
      <div className="flex justify-center items-center h-screen w-4/6">
        <div>
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
          <p className="mb-4">タスク一覧</p>
          <TodoList todoList={todoList} />
        </div>
      </div>
    </MainLayout>
  )
}

export default JotaiTest
