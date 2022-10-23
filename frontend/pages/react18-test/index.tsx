import { MainLayout } from '@/components/Layout/MainLayout'
import { Spinner } from '@/components/UI/Spiner'
import { useQueryComments } from '@/features/react18-test/hooks/useQueryComennts'
import { useQueryTasks } from '@/features/react18-test/hooks/useQueryTasks'
import { useQueryUsers } from '@/features/react18-test/hooks/useQueryUsers'
import React, { Suspense } from 'react'

const React18Test = () => {

  return (
    <MainLayout title="stripe">
      <div className="flex justify-center items-center h-screen w-4/6">
        <div>
          <Suspense fallback={<Spinner />}>
            <FetchUsers />
          </Suspense>
        </div>
      </div>
    </MainLayout>
  )
}

export const FetchUsers = () => {
  const { data } = useQueryUsers()
  //   if (status === 'loading') return <p>Loading...</p>
  //   if (status === 'error') return <p>Error</p>
  return (
    <div className="my-3 text-center">
      <p className="my-3 font-bold">User List</p>
      {data?.map((user) => (
        <p className="my-3 text-sm" key={user.id}>
          {user.username}
        </p>
      ))}
    </div>
  )
}

export const FetchComments = () => {
  const { data } = useQueryComments()

  return (
    <div className="my-3 text-center">
      <p className="my-3 font-bold">Comment List</p>
      {data?.map((comment) => (
        <p className="my-3 text-sm" key={comment.id}>
          {comment.name}
        </p>
      ))}
    </div>
  )
}

export const FetchTasks = () => {
  const { data } = useQueryTasks()

  return (
    <div className="my-3 text-center">
      <p className="my-3 font-bold">Task List</p>
      {data?.map((task) => (
        <p className="my-3 text-sm" key={task.id}>
          {task.title}
        </p>
      ))}
    </div>
  )
}

export default React18Test
