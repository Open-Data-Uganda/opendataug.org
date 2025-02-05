import { useMutation, useQueryClient } from '@tanstack/react-query';
import axios, { AxiosError } from 'axios';
import { backendUrl } from '../config';
import { useAuth } from '../context/AuthContext';

interface UseDeleteRequestProps {
  queryKey: string;
  url: string;
}

const useDeleteRequest = ({ queryKey, url }: UseDeleteRequestProps) => {
  const { userNumber, accessToken } = useAuth();

  const queryClient = useQueryClient();

  return useMutation<void, AxiosError>({
    mutationFn: async () => {
      await axios.delete(`${backendUrl}/${url}`, {
        withCredentials: true,
        headers: {
          'User-Number': userNumber,
          Authorization: `Bearer ${accessToken}`
        }
      });
    },

    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: [queryKey] });
    }
  });
};

export default useDeleteRequest;
