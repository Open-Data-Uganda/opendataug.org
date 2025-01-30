import { useMutation, useQueryClient } from '@tanstack/react-query';
import axios from 'axios';
import { backendUrl } from '../config';

interface PostRequestProps {
  url: string;
  queryKey: any;
}

const usePostRequest = ({ url, queryKey }: PostRequestProps) => {
  const token = localStorage.getItem('token');
  const userNumber = localStorage.getItem('userNumber');
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async (body) => {
      try {
        const response = await axios.post(`${backendUrl}/v1/${url}`, body, {
          headers: {
            Authorization: `Bearer ${token}`,
            'User-Number': userNumber
          }
        });
        return response.data;
      } catch (error) {
        if (axios.isAxiosError(error) && error.response) {
          throw new Error(error.response.data.message || 'An error occurred while processing your request.');
        } else {
          throw new Error('An error occurred while processing your request.');
        }
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: [queryKey] });
    },
    onError: (error) => {
      console.error(error);
    }
  });
};

export default usePostRequest;
