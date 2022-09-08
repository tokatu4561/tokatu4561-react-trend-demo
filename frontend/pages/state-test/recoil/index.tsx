import { MainLayout } from '@/components/Layout/MainLayout'
import { Todo } from '@/features/todo/types'
import { searchedTodoListSelector } from '@/stores/selectors/SearchedTodoListSelector'
import React from 'react'
import { useRecoilValue } from 'recoil'

const StripeTest = () => {
  const todoList = useRecoilValue(searchedTodoListSelector)

  return (
    <MainLayout title="stripe">
      <div>
        <p>タスク一覧</p>
        <ul>
          {todoList.map((todo: Todo, i: number) => {
            return <li key={`${todo.title}_${i}`}>{todo.title}</li>
          })}
        </ul>
      </div>
    </MainLayout>
  )
}

export default StripeTest
