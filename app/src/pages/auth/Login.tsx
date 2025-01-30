import { zodResolver } from '@hookform/resolvers/zod';
import React, { useEffect, useState } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Link, useNavigate } from 'react-router-dom';
import { z } from 'zod';
import Logo from '../../assets/logo.svg';
import Button from '../../components/Button';
import Input from '../../components/Input';
import { notifyError } from '../../components/toasts';
import { useAuth } from '../../context/AuthContext';
import { LoginSchema } from '../../types/schemas';

type Inputs = z.infer<typeof LoginSchema>;

const Login: React.FC = () => {
  const [disabled, setDisabled] = useState(false);
  const [loading, setLoading] = useState(false);

  const { login, userNumber, token, userRole } = useAuth();
  const {
    register,
    handleSubmit,
    watch,
    reset,
    formState: { errors }
  } = useForm<Inputs>({
    resolver: zodResolver(LoginSchema)
  });

  const emailValue = watch('email');
  const passwordValue = watch('password');
  const navigate = useNavigate();

  const onSubmit: SubmitHandler<Inputs> = async () => {
    setLoading(true);
    try {
      setDisabled(true);
      await login(emailValue, passwordValue);

      setDisabled(false);
      reset();
      navigate('/dashboard');
    } catch (error) {
      notifyError((error as Error).message);
      setDisabled(false);
    }
    setLoading(false);
  };

  useEffect(() => {
    if (userNumber && token && userRole) {
      navigate('/dashboard');
    }
  }, [token, userNumber, userRole, navigate]);

  return (
    <div className="flex min-h-screen">
      <div className="flex w-full items-center justify-center p-8 lg:w-1/2">
        <div className="w-full max-w-md">
          <img src={Logo} alt="Uganda Data Logo" className="mb-12 h-8" />

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
              <Link to="/signup" className="text-blue-600">
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
  );
};

export default Login;
