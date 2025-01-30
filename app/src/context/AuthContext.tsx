import { createContext, ReactNode, useContext, useEffect, useState } from 'react';
import { backendUrl } from '../config';

interface AuthContextProps {
  token: string | null;
  userNumber: string | null;
  userRole: string | null;
  login: (email: string, password: string) => Promise<void>;
  logout: () => void;
}

interface AuthProviderProps {
  children: ReactNode;
}

const AuthContext = createContext<AuthContextProps | undefined>(undefined);

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [token, setToken] = useState('');
  const [userNumber, setUserNumber] = useState('');
  const [userRole, setUserRole] = useState('');

  const login = async (email: string, password: string) => {
    try {
      const response = await fetch(`${backendUrl}/auth/login`, {
        method: 'POST',
        credentials: 'include',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ email, password })
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || 'Login failed');
      }

      const data = await response.json();

      setToken(data.token);
      setUserNumber(data.userNumber);
      setUserRole(data.role);
      localStorage.setItem('token', data.token);
      localStorage.setItem('userNumber', data.userNumber);
      localStorage.setItem('userRole', data.role);
    } catch (error) {
      throw error;
    }
  };

  const logout = () => {
    setToken('');
    setUserNumber('');
    setUserRole('');
    localStorage.removeItem('token');
    localStorage.removeItem('userRole');
    localStorage.removeItem('userNumber');
  };

  useEffect(() => {
    const storedToken = localStorage.getItem('token');
    const storedUserNumber = localStorage.getItem('userNumber');
    const storedUserRole = localStorage.getItem('userRole');

    if (storedToken && storedUserNumber && storedUserRole) {
      setToken(storedToken);
      setUserNumber(storedUserNumber);
      setUserRole(storedUserRole);
    }
  }, []);

  return <AuthContext.Provider value={{ token, userNumber, login, logout, userRole }}>{children}</AuthContext.Provider>;
};

export const useAuth = (): AuthContextProps => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};
