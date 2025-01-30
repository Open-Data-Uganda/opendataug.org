import { useQuery } from '@tanstack/react-query';
import axios from 'axios';
import { backendUrl } from '../config';

const useGetRequest = (initialUrl: string, queryKey: string, queryOptions?: any) => {
  const token = localStorage.getItem('token');
  const userNumber = localStorage.getItem('userNumber');

  const { data, isLoading, isPlaceholderData, isError } = useQuery({
    queryKey: [queryKey, queryOptions],
    staleTime: Infinity,
    queryFn: async () => {
      const { data } = await axios.get(`${backendUrl}/v1/${initialUrl}`, {
        headers: {
          Authorization: `Bearer ${token}`,
          'User-Number': userNumber
        }
      });

      return {
        data: data.data,
        total: data.total ?? null
      };
    }
  });

  return { data, isLoading, isError, isPlaceholderData };
};

export default useGetRequest;
