import { Todo } from '@/features/todo/types'
import { atom } from 'jotai'

export const todoListAtom = atom<Todo[]>([{ title: 'aaa' }])
