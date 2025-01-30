import { Suspense } from 'react';
import { Route, Routes } from 'react-router-dom';
import LoadingSpinner from './components/LoadingSpinner';
import { getRouteElement, routes } from './routes';

const App = () => {
  return (
    <Suspense fallback={<LoadingSpinner />}>
      <Routes>
        {routes.map((route) => (
          <Route key={route.path} path={route.path} element={getRouteElement(route)} />
        ))}
      </Routes>
    </Suspense>
  );
};

export default App;
