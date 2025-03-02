import { zodResolver } from '@hookform/resolvers/zod';
import React, { useEffect, useState } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Link, useNavigate } from 'react-router-dom';
import { z } from 'zod';

import Button from '../../components/Button';
import Input from '../../components/Input';
import SmallHeader from '../../components/SmallHeader';
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
    } catch (error) {
      notifyError((error as Error).message);
    } finally {
      setLoading(false);
      setDisabled(false);
    }
  };

  return (
    <div className="max-h-screen overflow-hidden">
      <SmallHeader />

      <div className="flex justify-center px-4 py-10">
        <div className="w-full max-w-md rounded-xl bg-white p-8">
          <div className="mb-8 text-center">
            <h1 className="mb-2 text-2xl font-bold text-gray-800">Access Uganda's Data</h1>
            <p className="text-sm text-gray-600">
              Get started with comprehensive data about Uganda's districts, villages, and administrative units 
              <span className="ml-1">🇺🇬</span>
            </p>
          </div>

          <form onSubmit={handleSubmit(onSubmit)} >
            <div className=" mb-6">
                <div className=" flex flex-col">
              <Input
                label="First name"
                type="text"
                 required
                error={errors.firstName?.message}
                {...register('firstName')}
              />

              <Input
                label="Last name"
                type="text"
                 required
                error={errors.lastName?.message}
                {...register('lastName')}
              />
            </div>

            <Input
              label="Email address"
              type="email"
               required
              error={errors.email?.message}
              {...register('email')}
            />

            </div>
            
            <div className="space-y-4">
              <label className="flex items-start mb-6">
                <input
                  type="checkbox"
                  className="mt-1 h-4 w-4 rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  checked={termsAccepted}
                  onChange={(e) => setTermsAccepted(e.target.checked)}
                  required
                />
                <span className="ml-2 text-xs text-gray-600">
                  By clicking Get started, you agree to our{' '}
                  <Link to="/terms" className="font-medium text-blue-600 hover:text-blue-500">
                    Terms of Service
                  </Link>{' '}
                  and{' '}
                  <Link to="/privacy" className="font-medium text-blue-600 hover:text-blue-500">
                    Privacy Policy
                  </Link>
                </span>
              </label>
            </div>

            <Button 
              type="submit" 
              disabled={disabled || !termsAccepted} 
              loading={loading} 
              fullWidth
              className="bg-blue-600 text-white hover:bg-blue-700"
            >
              Get started
            </Button>

            <p className="text-center mt-4 text-sm text-gray-600">
              Already have an account?{' '}
              <Link to="/login" className="font-medium text-blue-600 hover:text-blue-500">
                Login
              </Link>
            </p>
          </form>
        </div>
      </div>
    </div>
  );
};

export default SignUp;
