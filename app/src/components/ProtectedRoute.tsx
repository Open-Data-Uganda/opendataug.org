import { Navigate, Outlet } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';

export const ProtectedRoute = () => {
  const { accessToken, userNumber, userRole, isAuthenticated } = useAuth();

  if (!accessToken && !userNumber && !isAuthenticated && !userRole) {
    return <Navigate to="/" />;
  }

  return <Outlet />;
};
