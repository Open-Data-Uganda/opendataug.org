import { zodResolver } from '@hookform/resolvers/zod';
import React, { useEffect, useState } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Link, useNavigate } from 'react-router-dom';
import { z } from 'zod';

import Kids from '../../assets/kids.jpg';
import Button from '../../components/Button';
import Input from '../../components/Input';
import { notifyError, notifySuccess } from '../../components/toasts';
import { backendUrl } from '../../config';
import { useAuth } from '../../context/AuthContext';
import { SignUpSchema } from '../../types/schemas';

type Inputs = z.infer<typeof SignUpSchema>;

const SignUp: React.FC = () => {
  const [disabled, setDisabled] = useState(false);
  const [loading, setLoading] = useState(false);
  const [termsAccepted, setTermsAccepted] = useState(false);
  const { userNumber, isAuthenticated, userRole, isLoading, accessToken } = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    if (!isLoading && isAuthenticated && userNumber && userRole && accessToken) {
      navigate('/dashboard');
    }
  }, [isLoading, isAuthenticated, userNumber, userRole, navigate, accessToken]);

  const {
    register,
    handleSubmit,
    reset,
    formState: { errors }
  } = useForm<Inputs>({
    resolver: zodResolver(SignUpSchema)
  });

  console.log('backend url', backendUrl);

  const onSubmit: SubmitHandler<Inputs> = async (data) => {
    setLoading(true);
    setDisabled(true);

    try {
      const response = await fetch(`${backendUrl}/auth/register`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          first_name: data.firstName,
          other_name: data.lastName,
          email: data.email
        })
      });

      const result = await response.json();

      if (!response.ok) {
        throw new Error(result.message || 'Failed to sign up');
      }

      notifySuccess('Sign up successful! Please check your email to verify your account.');
      reset();
      navigate('/login');
    } catch (error) {
      notifyError((error as Error).message);
    } finally {
      setLoading(false);
      setDisabled(false);
    }
  };

  return (
    <div className=" h-screen">
      <div className="border-b border-gray-200">
        <div className="mx-auto flex w-full max-w-screen-2xl items-center justify-between px-6 py-4">
          <img
            src="https://tailwindui.com/plus/img/logos/mark.svg?color=indigo&shade=900"
            alt="Uganda Data Logo"
            className="h-8"
          />
        </div>
      </div>
      <div className="flex min-h-screen ">
        <div className="flex w-full items-center justify-center p-8 lg:w-1/2">
          <div className="w-full max-w-md">
            <div className="mb-8">
              <h1 className="mb-2 text-2xl font-bold">Access Uganda's Data</h1>
              <p className="text-gray-600">
                Get started with comprehensive data about Uganda's districts, villages, and administrative units 🇺🇬
              </p>
            </div>

            <form onSubmit={handleSubmit(onSubmit)}>
              <div className="grid lg:grid-cols-2 lg:gap-4">
                <Input
                  label="First name"
                  type="text"
                  placeholder="Enter your first name"
                  required
                  error={errors.firstName?.message}
                  {...register('firstName')}
                />

                <Input
                  label="Last name"
                  type="text"
                  placeholder="Enter your last name"
                  required
                  error={errors.lastName?.message}
                  {...register('lastName')}
                />
              </div>

              <Input
                label="Email address"
                type="email"
                placeholder="Email address"
                required
                error={errors.email?.message}
                {...register('email')}
              />

              <div className="mb-6 space-y-4">
                <label className="flex items-center">
                  <input
                    type="checkbox"
                    className="mr-2 rounded border-gray-300"
                    checked={termsAccepted}
                    onChange={(e) => setTermsAccepted(e.target.checked)}
                    required
                  />
                  <span className="text-sm">
                    By clicking Get started, you agree to our{' '}
                    <Link to="/terms" className="text-blue-600">
                      Terms of Service
                    </Link>{' '}
                    and{' '}
                    <Link to="/privacy" className="text-blue-600">
                      Privacy Policy
                    </Link>
                  </span>
                </label>
              </div>

              <Button type="submit" disabled={disabled || !termsAccepted} loading={loading} fullWidth>
                Get started
              </Button>

              <p className="mt-6 text-center">
                Already have an account?{' '}
                <Link to="/login" className="text-blue-600">
                  Login
                </Link>
              </p>
            </form>
          </div>
        </div>

        {/* Right side - Illustration and partners */}
        <div className="  flex-col justify-between bg-gray-50 p-8 lg:flex lg:w-1/2">
          <div className="flex flex-1 items-center justify-center">
            <img loading="lazy" src={Kids} alt="image of school kids" />
          </div>
        </div>
      </div>
    </div>
  );
};

export default SignUp;
