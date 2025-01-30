import { zodResolver } from '@hookform/resolvers/zod';
import React, { useState } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Link, useParams } from 'react-router-dom';
import { z } from 'zod';
import Logo from '../../assets/logo.svg';
import { ErrorMessage } from '../../components/ErrorMessage';
import { notifyError, notifySuccess } from '../../components/toasts';
import { backendUrl } from '../../config';
import { SetPasswordSchema } from '../../types/schemas';

type Inputs = z.infer<typeof SetPasswordSchema>;

const SetPassword: React.FC = () => {
  const [disabled, setDisabled] = useState(false);
  const [loading, setLoading] = useState(false);
  const { token } = useParams();

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
    <div className=" flex h-screen min-w-full items-center justify-center text-black">
      <div className="w-full p-4 sm:w-4/12 md:w-6/12 lg:w-3/12 ">
        <div className=" flex items-center justify-center">
          <img className=" mb-6 lg:mb-10 lg:h-20" src={Logo} alt="" />
        </div>

        <div className="mb-9">
          <h2 className=" text-2xl font-bold text-black sm:text-title-lg">Set Password</h2>

          <p>Set a password for your account.</p>
        </div>
        <form onSubmit={handleSubmit(onSubmit)}>
          <div className="mb-9">
            <label className="mb-2.5 block font-medium text-black" id="password">
              Password <span className=" text-red-500">*</span>
            </label>
            <div className="relative">
              <input id="password" type="password" {...register('password')} className="form-class" />
              {errors?.password && <ErrorMessage error={errors.password?.message} />}
            </div>
          </div>

          <div className="mb-9">
            <label className="mb-2.5 block font-medium text-black" id="confirm_password">
              Confirm Password <span className=" text-red-500">*</span>
            </label>
            <div className="relative">
              <input id="confirm_password" type="password" {...register('confirm_password')} className="form-class" />
              {errors?.confirm_password && <ErrorMessage error={errors.confirm_password?.message} />}
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
                'Set Password'
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

export default SetPassword;
