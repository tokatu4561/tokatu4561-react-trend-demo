import { Todo } from '@/features/todo/types'
import { atom } from 'recoil'

export const todoListState = atom<Todo[]>({
  key: 'todoList',
  default: [{ title: 'one' }, { title: 'two' }],
})
