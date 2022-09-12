import React, { FC } from 'react'
import { Todo } from '../types'

interface Props {
  todoList: Todo[]
}

export const TodoList: FC<Props> = ({ todoList }) => {
  return (
    <ul>
      {todoList.map((todo: Todo, i: number) => {
        return (
          <li className="p-2 mb-2 bg-white rounded" key={`${todo.title}_${i}`}>
            {todo.title}
          </li>
        )
      })}
    </ul>
  )
}
