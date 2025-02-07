import { zodResolver } from '@hookform/resolvers/zod';
import React, { useState } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Link, useParams } from 'react-router-dom';
import { z } from 'zod';
import Button from '../../components/Button';
import Input from '../../components/Input';
import SmallHeader from '../../components/SmallHeader';
import { notifyError, notifySuccess } from '../../components/toasts';
import { backendUrl } from '../../config';
import { ResetPasswordSchema } from '../../types/schemas';
type Inputs = z.infer<typeof ResetPasswordSchema>;

const ResetPassword: React.FC = () => {
  const [disabled, setDisabled] = useState(false);
  const [loading, setLoading] = useState(false);
  const {
    register,
    handleSubmit,
    watch,
    reset,
    formState: { errors }
  } = useForm<Inputs>({
    resolver: zodResolver(ResetPasswordSchema)
  });

  const token = useParams();

  const emailValue = watch('email');

  const onSubmit: SubmitHandler<Inputs> = async () => {
    setLoading(true);
    try {
      setDisabled(true);
      const response = await fetch(`${backendUrl}/auth/reset-password?token=${token}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          email: emailValue
        })
      });

      if (response.ok) {
        setDisabled(false);
        notifySuccess('A password reset link has been set to your email.');
        reset();
      } else {
        const errorData = await response.json();
        notifyError(errorData.message || 'An error occurred while resetting the password.');
      }
    } catch (error: any) {
      notifyError(error.message || 'An unexpected error occurred. Please try again.');
    } finally {
      setDisabled(false);
    }
    setLoading(false);
  };

  return (
    <div className="h-screen">
      <SmallHeader />
      <div className="flex min-h-screen justify-center">
        <div className="flex w-full items-center justify-center p-8 lg:w-1/2">
          <div className="w-full max-w-md">
            <div className="mb-8">
              <h1 className="mb-2 text-2xl font-bold">Reset Your Password</h1>
              <p className="text-gray-600">
                Enter your email address below and we'll send you instructions to reset your password.
              </p>
            </div>

            <form onSubmit={handleSubmit(onSubmit)}>
              <Input
                label="Email address"
                type="email"
                placeholder="Enter your registered email"
                required
                error={errors.email?.message}
                {...register('email')}
              />

              <div className="mb-6 mt-2">
                <p className="text-sm text-gray-500">
                  Make sure to use the email address associated with your account.
                </p>
              </div>

              <Button type="submit" disabled={disabled} loading={loading} fullWidth>
                Send Reset Instructions
              </Button>

              <div className="mt-6 space-y-4">
                <p className="text-center text-sm text-gray-600">
                  Remember your password?{' '}
                  <Link to="/login" className="text-blue-600 hover:text-blue-700">
                    Back to Login
                  </Link>
                </p>
                <p className="text-center text-sm text-gray-600">
                  Don't have an account?{' '}
                  <Link to="/" className="text-blue-600 hover:text-blue-700">
                    Sign Up
                  </Link>
                </p>
              </div>
            </form>

            <div className="mt-8">
              <div className="rounded-md bg-blue-50 p-4">
                <div className="flex">
                  <div className="flex-shrink-0">
                    <svg
                      className="h-5 w-5 text-blue-400"
                      xmlns="http://www.w3.org/2000/svg"
                      viewBox="0 0 20 20"
                      fill="currentColor">
                      <path
                        fillRule="evenodd"
                        d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"
                        clipRule="evenodd"
                      />
                    </svg>
                  </div>
                  <div className="ml-3">
                    <p className="text-sm text-blue-700">
                      After submitting, please check your email inbox and spam folder. The reset link will expire in 1
                      hour.
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ResetPassword;
