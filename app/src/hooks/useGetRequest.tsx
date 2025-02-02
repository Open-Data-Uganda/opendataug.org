import { useQuery } from '@tanstack/react-query';
import axios from 'axios';
import { backendUrl } from '../config';
import { useAuth } from '../context/AuthContext';

interface GetRequestProps {
  url: string;
  queryKey: string | string[];
  params?: Record<string, any>;
}

const useGetRequest = ({ url, queryKey, params }: GetRequestProps) => {
  const { userNumber, accessToken } = useAuth();

  return useQuery({
    queryKey: Array.isArray(queryKey) ? queryKey : [queryKey],
    queryFn: async () => {
      try {
        const response = await axios.get(`${backendUrl}/${url}`, {
          params,
          withCredentials: true,
          headers: {
            'Content-Type': 'application/json',
            'User-Number': userNumber,
            Authorization: `Bearer ${accessToken}`
          }
        });

        if (response.data.data && response.data.total !== undefined) {
          return {
            data: response.data.data,
            total: response.data.total
          };
        }

        return response.data;
      } catch (error) {
        if (axios.isAxiosError(error) && error.response) {
          throw new Error(error.response.data.message || 'An error occurred while fetching data.');
        } else {
          throw new Error('An error occurred while fetching data.');
        }
      }
    }
  });
};

export default useGetRequest;
