import { zodResolver } from '@hookform/resolvers/zod';
import React, { useState } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Link, useNavigate, useSearchParams } from 'react-router-dom';
import { z } from 'zod';
import Button from '../../components/Button';
import Input from '../../components/Input';
import { notifyError, notifySuccess } from '../../components/toasts';
import { backendUrl } from '../../config';
import { SetPasswordSchema } from '../../types/schemas';

type Inputs = z.infer<typeof SetPasswordSchema>;

const SetPassword: React.FC = () => {
  const navigate = useNavigate();
  const [disabled, setDisabled] = useState(false);
  const [loading, setLoading] = useState(false);
  const [searchParams] = useSearchParams();
  const token = searchParams.get('token');

  const {
    register,
    handleSubmit,
    reset,
    formState: { errors }
  } = useForm<Inputs>({
    resolver: zodResolver(SetPasswordSchema)
  });

  const onSubmit: SubmitHandler<Inputs> = async (data: any) => {
    setLoading(true);
    try {
      setDisabled(true);
      const response = await fetch(`${backendUrl}/auth/set-password?token=${token}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          password: data.password,
          confirm_password: data.confirm_password
        })
      });

      if (response.ok) {
        setDisabled(false);
        notifySuccess('A password has been set. You can now sign in.');
        reset();
        navigate('/login');
      } else {
        const errorData = await response.json();
        notifyError(errorData.message || 'An error occurred while setting the password.');
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
              <h1 className="mb-2 text-2xl font-bold">Create New Password</h1>
              <p className="text-gray-600">
                Please choose a strong password to secure your account. Make sure it's at least 8 characters long.
              </p>
            </div>

            <form onSubmit={handleSubmit(onSubmit)}>
              <Input
                label="New Password"
                type="password"
                placeholder="Enter your new password"
                required
                error={errors.password?.message}
                {...register('password')}
              />

              <div className="mb-4 mt-2">
                <ul className="space-y-1 text-sm text-gray-600">
                  <li className="flex items-center">
                    <svg className="mr-2 h-4 w-4 text-green-500" fill="currentColor" viewBox="0 0 20 20">
                      <path
                        fillRule="evenodd"
                        d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                        clipRule="evenodd"
                      />
                    </svg>
                    At least 8 characters long
                  </li>
                  <li className="flex items-center">
                    <svg className="mr-2 h-4 w-4 text-green-500" fill="currentColor" viewBox="0 0 20 20">
                      <path
                        fillRule="evenodd"
                        d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                        clipRule="evenodd"
                      />
                    </svg>
                    Include numbers and special characters
                  </li>
                </ul>
              </div>

              <Input
                label="Confirm Password"
                type="password"
                placeholder="Confirm your new password"
                required
                error={errors.confirm_password?.message}
                {...register('confirm_password')}
              />

              <div className="mt-6">
                <Button type="submit" disabled={disabled} loading={loading} fullWidth>
                  Set New Password
                </Button>
              </div>

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
                        After setting your new password, you'll be redirected to the login page.
                      </p>
                    </div>
                  </div>
                </div>
              </div>

              <div className="mt-6 text-center">
                <p className="text-sm text-gray-600">
                  Need help?{' '}
                  <Link to="/contact" className="text-blue-600 hover:text-blue-700">
                    Contact Support
                  </Link>
                </p>
              </div>
            </form>
          </div>
        </div>

        <div className="hidden flex-col justify-between bg-gray-50 p-8 lg:flex lg:w-1/2">
          <div className="flex flex-1 items-center justify-center">
            <div className="text-center">
              <h2 className="text-2xl font-bold text-gray-900">Secure Your Account</h2>
              <p className="mt-4 text-lg text-gray-600">
                Creating a strong password is an important step in protecting your data and maintaining secure access to
                Uganda's comprehensive database.
              </p>
            </div>
          </div>

          <div className="text-center">
            <p className="mb-4 text-gray-600">Trusted by organizations across Uganda</p>
            <div className="flex items-center justify-center space-x-8">{/* Partner logos remain unchanged */}</div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default SetPassword;
