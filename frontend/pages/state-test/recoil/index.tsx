import { MainLayout } from '@/components/Layout/MainLayout'
import CreateForm from '@/features/todo/components/CreateForm'
import SearchForm from '@/features/todo/components/SearchForm'
import { Todo } from '@/features/todo/types'
import { searchedTodoListSelector } from '@/stores/selectors/SearchedTodoListSelector'
import React from 'react'
import { useRecoilValue } from 'recoil'

const StripeTest = () => {
  const todoList = useRecoilValue(searchedTodoListSelector)

  return (
    <MainLayout title="stripe">
      <div className="flex justify-center items-center h-screen w-4/6">
        <div>
          <CreateForm />
          <SearchForm />
          <p className="mb-4">タスク一覧</p>
          <ul>
            {todoList.map((todo: Todo, i: number) => {
              return (
                <li
                  className="p-2 mb-2 bg-white rounded"
                  key={`${todo.title}_${i}`}
                >
                  {todo.title}
                </li>
              )
            })}
          </ul>
        </div>
      </div>
    </MainLayout>
  )
}

export default StripeTest
