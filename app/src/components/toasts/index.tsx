import { toast } from 'react-toastify';

export function notifySuccess(message: string) {
  toast.success(message, {
    position: 'bottom-right',
    hideProgressBar: true,
    autoClose: 2500
  });
}

export function notifyError(message: string) {
  toast.error(message, {
    position: 'bottom-right',
    hideProgressBar: true,
    autoClose: 2500
  });
}
