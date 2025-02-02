import { createContext, ReactNode, useContext, useEffect, useState } from 'react';
import { notifyError } from '../components/toasts';
import { backendUrl } from '../config';

interface AuthContextProps {
  isAuthenticated: boolean;
  userNumber: string | null;
  userRole: string | null;
  accessToken: string | null;
  isLoading: boolean;
  login: (email: string, password: string) => Promise<void>;
  logout: () => Promise<void>;
  checkAuthStatus: () => Promise<void>;
}

interface AuthProviderProps {
  children: ReactNode;
}

const AuthContext = createContext<AuthContextProps | undefined>(undefined);

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [userNumber, setUserNumber] = useState<string | null>(null);
  const [userRole, setUserRole] = useState<string | null>(null);
  const [accessToken, setAccessToken] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  const checkAuthStatus = async () => {
    try {
      setIsLoading(true);
      const response = await fetch(`${backendUrl}/auth/refresh`, {
        method: 'POST',
        credentials: 'include',
        headers: {
          'Content-Type': 'application/json'
        }
      });

      if (!response.ok) {
        setIsAuthenticated(false);
        setUserNumber(null);
        setUserRole(null);
        setAccessToken(null);
        return;
      }

      const data = await response.json();
      setIsAuthenticated(true);
      setUserNumber(data.user_number);
      setUserRole(data.role);
      setAccessToken(data.access_token);
    } catch (error) {
      console.error('Auth check failed:', error);
      setIsAuthenticated(false);
      setUserNumber(null);
      setUserRole(null);
      setAccessToken(null);
    } finally {
      setIsLoading(false);
    }
  };

  const login = async (email: string, password: string) => {
    try {
      setIsLoading(true);
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
      setIsAuthenticated(true);
      setUserNumber(data.user_number);
      setUserRole(data.role);
      setAccessToken(data.access_token);
    } catch (error) {
      notifyError('An error occurred');
      throw error;
    } finally {
      setIsLoading(false);
    }
  };

  const logout = async () => {
    try {
      setIsLoading(true);
      await fetch(`${backendUrl}/auth/logout`, {
        method: 'POST',
        credentials: 'include'
      });
    } finally {
      setIsAuthenticated(false);
      setUserNumber(null);
      setUserRole(null);
      setAccessToken(null);
      setIsLoading(false);
    }
  };

  useEffect(() => {
    checkAuthStatus();

    const refreshInterval = setInterval(
      () => {
        checkAuthStatus();
      },
      4 * 60 * 1000
    );

    const handleVisibilityChange = () => {
      if (document.visibilityState === 'visible') {
        checkAuthStatus();
      }
    };
    document.addEventListener('visibilitychange', handleVisibilityChange);

    return () => {
      clearInterval(refreshInterval);
      document.removeEventListener('visibilitychange', handleVisibilityChange);
    };
  }, []);

  return (
    <AuthContext.Provider
      value={{
        isAuthenticated,
        userNumber,
        userRole,
        accessToken,
        isLoading,
        login,
        logout,
        checkAuthStatus
      }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = (): AuthContextProps => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};
