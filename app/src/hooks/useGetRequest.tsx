import { useQuery } from '@tanstack/react-query';
import axios from 'axios';
import { notifyError } from '../components/toasts';
import { backendUrl } from '../config';
import { useAuth } from '../context/AuthContext';

interface GetRequestProps {
  url: string;
  queryKey: string | string[];
  params?: Record<string, any>;
}

export const useGetRequest = ({ url, queryKey, params }: GetRequestProps) => {
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
            Authorization: `Bearer ${accessToken}`,
          },
        });

        if (response.data.data && response.data.total !== undefined) {
          return {
            data: response.data.data,
            total: response.data.total,
          };
        }

        return response.data;
      } catch (error) {
        if (axios.isAxiosError(error) && error.response) {
          notifyError('An error occurred while fetching data');
        } else {
          notifyError('An error occurred while fetching data');
        }
      }
    },
  });
};
