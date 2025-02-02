import { useMutation, useQueryClient } from '@tanstack/react-query';
import axios from 'axios';
import { backendUrl } from '../config';
import { useAuth } from '../context/AuthContext';

interface PostRequestProps {
  url: string;
  queryKey: string;
}

const usePostRequest = ({ url, queryKey }: PostRequestProps) => {
  const queryClient = useQueryClient();
  const { userNumber, accessToken } = useAuth();

  return useMutation({
    mutationFn: async (body) => {
      try {
        const response = await axios.post(`${backendUrl}/${url}`, body, {
          withCredentials: true,
          headers: {
            'Content-Type': 'application/json',
            'User-Number': userNumber,
            Authorization: `Bearer ${accessToken}`
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
