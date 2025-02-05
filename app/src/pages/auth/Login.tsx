import { zodResolver } from '@hookform/resolvers/zod';
import React, { useEffect, useState } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Link, useNavigate } from 'react-router-dom';
import { z } from 'zod';
import Button from '../../components/Button';
import Input from '../../components/Input';
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
      <div className="border-b border-gray-200">
        <div className="mx-auto flex w-full max-w-screen-2xl items-center justify-between px-6 py-4">
          <img
            src="https://tailwindui.com/plus/img/logos/mark.svg?color=indigo&shade=900"
            alt="Uganda Data Logo"
            className="h-8"
          />
        </div>
      </div>

      <div className="flex min-h-screen">
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

        <div className="hidden flex-col justify-between bg-gray-50 p-8 lg:flex lg:w-1/2">
          <div className="flex flex-1 items-center justify-center">
            {/* Add illustration showing Uganda map or data visualization */}
          </div>

          <div className="text-center">
            <p className="mb-4 text-gray-600">Trusted by organizations across Uganda</p>
            <div className="flex items-center justify-center space-x-8">{/* Add Ugandan partner logos here */}</div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Login;
