import { useMutation, useQueryClient } from '@tanstack/react-query';
import axios from 'axios';
import { backendUrl } from '../config';

const useDeleteRequest = (queryKey: string, url: string) => {
  const token = localStorage.getItem('token');
  const userIdentifier = localStorage.getItem('userNumber');

  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async () => {
      await axios.delete(`${backendUrl}/v1/${url}`, {
        headers: {
          Authorization: `Bearer ${token}`,
          'User-Number': userIdentifier
        }
      });
    },

    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: [queryKey] });
    }
  });
};

export default useDeleteRequest;
