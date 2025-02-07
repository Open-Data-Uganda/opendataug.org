import { useEffect, useState } from 'react';
import { Route, Routes, useLocation } from 'react-router-dom';
import LoadingSpinner from './components/LoadingSpinner';
import PageTitle from './components/PageTitle';
import { ProtectedRoute } from './components/ProtectedRoute';
import { FallbackProvider } from './composables/FallbackProvider';
import Overview from './features/apiKeys';
import Login from './pages/auth/Login';
import ResetPassword from './pages/auth/ResetPassword';
import SetPassword from './pages/auth/SetPassword';
import SignUp from './pages/auth/SignUp';
import NotFound from './pages/NotFound';
import PageNotFound from './pages/PageNotFound';

const App = () => {
  const [loading, setLoading] = useState(true);

  const { pathname } = useLocation();

  useEffect(() => {
    window.scrollTo(0, 0);
  }, [pathname]);

  useEffect(() => {
    setTimeout(() => setLoading(false), 50);
  }, []);

  return loading ? (
    <div className="flex h-screen items-center justify-center">
      <LoadingSpinner />
    </div>
  ) : (
    <FallbackProvider>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/" element={<SignUp />} />
        <Route path="/reset-password" element={<ResetPassword />} />
        <Route
          path="/set-password/:token"
          element={
            <>
              <PageTitle title="Set Password | Uganda Open Data" />
              <SetPassword />
            </>
          }
        />
        <Route path="*" element={<PageNotFound />} />

        <Route element={<ProtectedRoute />}>
          <Route path="/dashboard" element={<Overview />} />
          <Route path="/dashboard/*" element={<NotFound />} />
        </Route>
      </Routes>
    </FallbackProvider>
  );
};

export default App;
