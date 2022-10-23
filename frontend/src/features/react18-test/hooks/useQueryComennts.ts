import axios from 'axios'
import { useQuery } from 'react-query'

const getComments = async () => {
  const { data } = await axios
    .get('https://jsonplaceholder.typicode.com/comments?_limit=3')
    .then()
  return data
}

export const useQueryComments = () => {
  return useQuery({
    queryKey: ['comments'],
    queryFn: getComments,
    staleTime: Infinity,
  })
}