import { FC } from 'react';

interface ErrorMessageProps {
  error: any;
}

export const ErrorMessage: FC<ErrorMessageProps> = ({ error }) => {
  return error && <p className="mt-2 text-sm text-red-400">{error}</p>;
};
