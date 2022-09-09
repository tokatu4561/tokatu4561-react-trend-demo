import { todoTitleFormState } from '@/stores/atoms/searchTextFormAtom'
import { todoListState } from '@/stores/atoms/todoListAtom'
import React, { useCallback } from 'react'
import { useRecoilValue, useSetRecoilState } from 'recoil'

const CreateForm = () => {
  const todoTitleFormValue = useRecoilValue(todoTitleFormState)
  const todoList = useRecoilValue(todoListState)
  const setTodoTitleFormValue = useSetRecoilState(todoTitleFormState)
  const setTodoList = useSetRecoilState(todoListState)

  const onChange = useCallback(
    (event: React.ChangeEvent<HTMLInputElement>) => {
      // 先に取得したsetTodoTitleFormValueに対して更新したい値を渡して実行
      setTodoTitleFormValue(event.target.value)
    },
    [setTodoTitleFormValue]
  )

  const onClick = () => {
    setTodoList([...todoList, { title: todoTitleFormValue }])
    setTodoTitleFormValue('')
  }

  return (
    <>
      <label>
        タスク名:
        <input
          type="text"
          value={todoTitleFormValue}
          onChange={onChange}
          name="title"
          style={{ margin: 12 }}
        />
      </label>
      <button onClick={onClick}>追加</button>
    </>
  )
}

export default CreateForm
