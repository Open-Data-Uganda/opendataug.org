import { zodResolver } from '@hookform/resolvers/zod';
import React, { useEffect, useState } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Link, useNavigate } from 'react-router-dom';
import { z } from 'zod';

import Button from '../../components/Button';
import Input from '../../components/Input';
import SmallHeader from '../../components/SmallHeader';
import { notifyError } from '../../components/toasts';
import { useAuth } from '../../context/AuthContext';
import { LoginSchema } from '../../types/schemas';

type Inputs = z.infer<typeof LoginSchema>;

const Login: React.FC = () => {
  const [disabled, setDisabled] = useState(false);
  const [loading, setLoading] = useState(false);
  const { login, userNumber, isAuthenticated, userRole, isLoading, accessToken } = useAuth();
  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm<Inputs>({
    resolver: zodResolver(LoginSchema),
  });

  const navigate = useNavigate();

  useEffect(() => {
    if (!isLoading && isAuthenticated && userNumber && userRole && accessToken) {
      navigate('/dashboard');
    }
  }, [isLoading, isAuthenticated, userNumber, userRole, navigate, accessToken]);

  const onSubmit: SubmitHandler<Inputs> = async data => {
    setLoading(true);
    setDisabled(true);
    try {
      await login(data.email, data.password);
    } catch {
      notifyError('Login failed. Please check your credentials and try again.');
    } finally {
      reset();
      setLoading(false);
      setDisabled(false);
    }
  };

  return (
    <div className="max-h-screen overflow-hidden">
      <SmallHeader />

      <div className="mx-auto flex  justify-center px-4 py-10">
        <div className="flex w-full items-center justify-center p-8 lg:w-1/2">
          <div className="w-full max-w-md bg-white p-8">
            <div className="mb-8 text-center">
              <h1 className="mb-2 text-2xl font-bold text-gray-800">Welcome back</h1>
              <p className="text-gray-600">Access Uganda's comprehensive data through our simple API 🇺🇬</p>
            </div>

            <form onSubmit={handleSubmit(onSubmit)}>
              <Input
                label="Email address"
                type="email"
                autoComplete="email"
                required
                error={errors.email?.message}
                {...register('email')}
              />

              <Input
                label="Password"
                type="password"
                required
                error={errors.password?.message}
                {...register('password')}
              />

              <Link
                to="/reset-password"
                className="mb-6 block text-sm text-blue-600 transition-colors hover:text-blue-800">
                Forgot your password?
              </Link>

              <Button
                type="submit"
                disabled={disabled}
                loading={loading}
                fullWidth
                className="bg-blue-600 transition-colors hover:bg-blue-700">
                Login
              </Button>

              <p className="mt-6 text-center text-sm text-gray-600">
                Don't have an account?{' '}
                <Link to="/" className="text-sm font-medium text-blue-600 transition-colors hover:text-blue-500">
                  Sign Up
                </Link>
              </p>
            </form>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Login;
