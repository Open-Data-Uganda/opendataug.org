import { useMutation, useQueryClient } from '@tanstack/react-query';
import axios from 'axios';
import { backendUrl } from '../config';

interface PatchRequestProps {
  url: string;
  queryKey: string;
}

const usePatchRequest = ({ url, queryKey }: PatchRequestProps) => {
  const token = localStorage.getItem('token');
  const userIdentifier = localStorage.getItem('userNumber');
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (body) => {
      return axios.patch(`${backendUrl}/v1/${url}`, body, {
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

export default usePatchRequest;
