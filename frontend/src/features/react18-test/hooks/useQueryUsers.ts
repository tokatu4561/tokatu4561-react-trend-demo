import { delay } from '@/lib/utils'
import axios from 'axios'
import { useQuery } from 'react-query'

type GetUsersType = {
    data: any
}

const getUsers = async () => {
  const { data } = await axios
    .get<GetUsersType>('https://jsonplaceholder.typicode.com/users?_limit=3')
    .then()
  return data
}

export const useQueryUsers = () => {
  return useQuery({
    queryKey: ['users'],
    queryFn: getUsers,
    staleTime: Infinity,
  })
}