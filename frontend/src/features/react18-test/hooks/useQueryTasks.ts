import axios from 'axios'
import { useQuery } from 'react-query'

const getTasks = async () => {
  const { data } = await axios
    .get('https://jsonplaceholder.typicode.com/todos?_limit=3')
    .then()
  return data
}

export const useQueryTasks = () => {
  return useQuery({
    queryKey: ['tasks'],
    queryFn: getTasks,
    staleTime: Infinity,
  })
}