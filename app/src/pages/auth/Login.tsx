import { zodResolver } from '@hookform/resolvers/zod';
import React, { useEffect, useState } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Link, useNavigate } from 'react-router-dom';
import { z } from 'zod';

import Button from '../../components/Button';
import Input from '../../components/Input';
import SmallHeader from '../../components/SmallHeader';
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
    formState: { errors }
  } = useForm<Inputs>({
    resolver: zodResolver(LoginSchema)
  });

  const navigate = useNavigate();

  useEffect(() => {
    if (!isLoading && isAuthenticated && userNumber && userRole && accessToken) {
      navigate('/dashboard');
    }
  }, [isLoading, isAuthenticated, userNumber, userRole, navigate, accessToken]);

  const onSubmit: SubmitHandler<Inputs> = async (data) => {
    setLoading(true);
    setDisabled(true);
    try {
      await login(data.email, data.password);
    } catch (error) {
    } finally {
      reset();
      setLoading(false);
      setDisabled(false);
    }
  };

  return (
    <div className="h-screen">
      <SmallHeader />

      <div className="mx-auto flex min-h-screen justify-center">
        <div className="flex w-full items-center justify-center p-8 lg:w-1/2">
          <div className="w-full max-w-md">
            <div className="mb-8">
              <h1 className="mb-2 text-2xl font-bold">Welcome back</h1>
              <p className="text-gray-600">Access Uganda's comprehensive data through our simple API 🇺🇬</p>
            </div>

            <form onSubmit={handleSubmit(onSubmit)}>
              <Input
                label="Email address"
                type="email"
                placeholder="Email address"
                required
                error={errors.email?.message}
                {...register('email')}
              />

              <Input
                label="Password"
                type="password"
                placeholder="Password"
                required
                error={errors.password?.message}
                {...register('password')}
              />

              <Link to="/reset-password" className="mb-6 block text-blue-600">
                Forgot your password?
              </Link>

              <Button type="submit" disabled={disabled} loading={loading} fullWidth>
                Login
              </Button>

              <p className="mt-6 text-center">
                Don't have an account?{' '}
                <Link to="/" className="text-blue-600">
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
