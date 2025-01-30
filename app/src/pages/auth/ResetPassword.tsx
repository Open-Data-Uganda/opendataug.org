import { zodResolver } from '@hookform/resolvers/zod';
import React, { useState } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Link, useParams } from 'react-router-dom';
import { z } from 'zod';
import Logo from '../../assets/logo.svg';
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
    <div className=" flex h-screen min-w-full items-center justify-center text-black">
      <div className="w-full p-4 sm:w-4/12 md:w-6/12 lg:w-3/12 ">
        <div className=" flex items-center justify-center">
          <img className=" mb-6 lg:mb-10 lg:h-20" src={Logo} alt="" />
        </div>

        <div className="mb-9">
          <h2 className=" text-2xl font-bold text-black sm:text-title-xl2">Reset Password</h2>

          <p>We will send a password reset link to your email.</p>
        </div>
        <form onSubmit={handleSubmit(onSubmit)}>
          <div className="mb-9">
            <label className="mb-2.5 block font-medium text-black" id="email">
              Email <span className=" text-red-500">*</span>
            </label>
            <div className="relative">
              <input id="email" type="email" {...register('email')} placeholder="Email" className="form-class" />
              {errors.email && <p className="text-red-500">{errors.email.message}</p>}
            </div>
          </div>

          <div className="mb-5">
            <button
              disabled={disabled}
              type="submit"
              className="inline-flex w-full items-center justify-center rounded bg-primary py-2 text-center font-medium text-white hover:bg-opacity-90">
              {loading ? (
                <svg
                  className="-ml-1 mr-3 h-5 w-5 animate-spin text-black"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24">
                  <circle className="opacity-25" stroke="currentColor" strokeWidth="4" cx="12" cy="12" r="10"></circle>
                  <path
                    className="opacity-75"
                    fill="#fff"
                    d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
              ) : (
                'Send Email'
              )}
            </button>
          </div>

          <div className="mt-6 text-center">
            <p>
              Already have an account?{' '}
              <Link to="/" className="text-primary underline">
                Sign In
              </Link>
            </p>
          </div>
        </form>
      </div>
    </div>
  );
};

export default ResetPassword;
