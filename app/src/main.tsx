import { ThemeProvider } from '@emotion/react';
import { createTheme } from '@mui/material';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { ReactQueryDevtools } from '@tanstack/react-query-devtools';
import React, { ReactNode } from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter as Router } from 'react-router-dom';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import App from './App';
import { AuthProvider } from './context/AuthContext';
import './css/style.css';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
      staleTime: Infinity,
      retry: 2
    }
  }
});

const customTheme = createTheme({
  typography: {}
});

const Providers = ({ children }: { children: ReactNode }) => {
  return (
    <AuthProvider>
      <Router>
        <QueryClientProvider client={queryClient}>
          <ThemeProvider theme={customTheme}>
            {children}
            <ToastContainer />
            <ReactQueryDevtools />
          </ThemeProvider>
        </QueryClientProvider>
      </Router>
    </AuthProvider>
  );
};

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <Providers>
      <App />
    </Providers>
  </React.StrictMode>
);
