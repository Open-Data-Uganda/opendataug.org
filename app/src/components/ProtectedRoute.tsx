import { PropsWithChildren } from 'react';
import { Navigate } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';

export const ProtectedRoute = ({ children }: PropsWithChildren) => {
  const { token, userNumber } = useAuth();

  if (!token && !userNumber) {
    return <Navigate to="/" />;
  }

  return <>{children}</>;
};
