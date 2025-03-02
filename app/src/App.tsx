import { useEffect, useState } from 'react';
import { Route, Routes, useLocation } from 'react-router-dom';
import LoadingSpinner from './components/LoadingSpinner';
import PageTitle from './components/PageTitle';
import { ProtectedRoute } from './components/ProtectedRoute';
import { FallbackProvider } from './composables/FallbackProvider';
import Overview from './features/apiKeys';
import EditProfile from './features/EditProfile';
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
        <Route
          path="/login"
          element={
            <>
              <PageTitle title="Login | Open Data Uganda" />
              <Login />
            </>
          }
        />
        <Route
          path="/"
          element={
            <>
              <PageTitle title="Sign Up | Open Data Uganda" />
              <SignUp />
            </>
          }
        />
        <Route
          path="/reset-password"
          element={
            <>
              <PageTitle title="Reset Password | Open Data Uganda" />
              <ResetPassword />
            </>
          }
        />
        <Route
          path="/set-password"
          element={
            <>
              <PageTitle title="Set Password | Open Data Uganda" />
              <SetPassword />
            </>
          }
        />
        <Route
          path="*"
          element={
            <>
              <PageTitle title="Page Not Found | Open Data Uganda" />
              <PageNotFound />
            </>
          }
        />

        <Route element={<ProtectedRoute />}>
          <Route
            path="/dashboard"
            element={
              <>
                <PageTitle title="Dashboard | Open Data Uganda" />
                <Overview />
              </>
            }
          />
          <Route
            path="/dashboard/settings"
            element={
              <>
                <PageTitle title="Profile | Open Data Uganda" />
                <EditProfile />
              </>
            }
          />

          <Route
            path="/dashboard/*"
            element={
              <>
                <PageTitle title="Not Found | Open Data Uganda" />
                <NotFound />
              </>
            }
          />
        </Route>
      </Routes>
    </FallbackProvider>
  );
};

export default App;
