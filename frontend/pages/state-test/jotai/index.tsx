import { MainLayout } from '@/components/Layout/MainLayout'
import { TodoList } from '@/features/todo/components/TodoList'
import { todoListAtom } from '@/stores/atoms/jotai/todoListAtim'

import { useAtom } from 'jotai'
import React from 'react'

const JotaiTest = () => {
  const [todoList, setTodoList] = useAtom(todoListAtom)

  return (
    <MainLayout title="stripe">
      <div className="flex justify-center items-center h-screen w-4/6">
        <div>
          <p className="mb-4">タスク一覧</p>
          <TodoList todoList={todoList} />
        </div>
      </div>
    </MainLayout>
  )
}

export default JotaiTest
