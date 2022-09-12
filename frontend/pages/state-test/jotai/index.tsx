import { MainLayout } from '@/components/Layout/MainLayout'
import CreateForm from '@/features/todo/components/CreateForm'
import SearchForm from '@/features/todo/components/SearchForm'
import { TodoList } from '@/features/todo/components/TodoList'
import { Todo } from '@/features/todo/types'

import { useAtom, atom } from 'jotai'
import React from 'react'

const todoListAtom = atom<Todo[]>([{ title: 'aaa' }])
const searchedTodoListAtom = atom((get) =>
  get(todoListAtom).filter((todo) => todo.title)
)

const StripeTest = () => {
  const [todoList, setTodoList] = useAtom(searchedTodoListAtom)

  return (
    <MainLayout title="stripe">
      <div className="flex justify-center items-center h-screen w-4/6">
        <div>
          <CreateForm />
          <SearchForm />
          <p className="mb-4">タスク一覧</p>
          <TodoList todoList={todoList} />
        </div>
      </div>
    </MainLayout>
  )
}

export default StripeTest
