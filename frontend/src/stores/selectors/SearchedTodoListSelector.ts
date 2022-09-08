import { Todo } from '@/features/todo/types'
import { selector } from 'recoil'
import { todoListState } from '../atoms/todoListAtom'
import { searchTextFormState } from '../atoms/todoTitleFormAtom'

export const searchedTodoListSelector = selector<Todo[]>({
  key: 'searchedTodoListSelector',
  get: ({ get }) => {
    const todoList: Todo[] = get(todoListState)

    const searchText: string = get(searchTextFormState)

    return searchText
      ? todoList.filter((t) => t.title.includes(searchText))
      : todoList
  },
})
