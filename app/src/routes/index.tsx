import { lazy } from 'react';
import PageTitle from '../components/PageTitle';
import { ProtectedRoute } from '../components/ProtectedRoute';
import PageNotFound from '../pages/PageNotFound';

const Login = lazy(() => import('../pages/auth/Login'));
const SignUp = lazy(() => import('../pages/auth/SignUp'));
const ResetPassword = lazy(() => import('../pages/auth/ResetPassword'));
const SetPassword = lazy(() => import('../pages/auth/SetPassword'));
const APIKeys = lazy(() => import('../features/apiKeys'));
const NotFound = lazy(() => import('../pages/NotFound'));

interface RouteConfig {
  path: string;
  element: React.ReactNode;
  title: string;
  protected?: boolean;
}

export const routes: RouteConfig[] = [
  {
    path: '/login',
    element: <Login />,
    title: 'Login | Open Data Uganda'
  },
  {
    path: '/',
    element: <SignUp />,
    title: 'Signup | Open Data Uganda'
  },
  {
    path: '/reset-password',
    element: <ResetPassword />,
    title: 'Reset Password | Open Data Uganda'
  },
  {
    path: '/set-password/:token',
    element: <SetPassword />,
    title: 'Set Password | Open Data Uganda'
  },
  {
    path: '/dashboard',
    element: <APIKeys />,
    title: 'API Keys | Open Data Uganda',
    protected: true
  },
  {
    path: '/dashboard/*',
    element: <NotFound />,
    title: 'Page Not Found | Open Data Uganda',
    protected: true
  },
  {
    path: '*',
    element: <PageNotFound />,
    title: 'Page Not Found | Open Data Uganda'
  }
];

export const getRouteElement = (route: RouteConfig) => {
  const Element = (
    <>
      <PageTitle title={route.title} />
      {route.element}
    </>
  );

  if (route.protected) {
    return <ProtectedRoute>{Element}</ProtectedRoute>;
  }

  return Element;
};
