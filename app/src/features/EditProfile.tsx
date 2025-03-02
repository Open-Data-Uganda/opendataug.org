import { zodResolver } from '@hookform/resolvers/zod';
import { useState } from 'react';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import Button from '../components/Button';
import Input from '../components/Input';
import { DeleteModal } from '../components/Modals/DeleteModal';
import { notifyError, notifySuccess } from '../components/toasts';
import useDeleteRequest from '../hooks/useDeleteRequest';
import useGetRequest from '../hooks/useGetRequest';
import usePatchRequest from '../hooks/usePatchRequest';
import DefaultLayout from '../layout/DefaultLayout';
import { EditProfileSchema } from '../types/schemas';

type Inputs = z.infer<typeof EditProfileSchema>;

const Settings = () => {
  const [loading, setLoading] = useState(false);
  const [showDeleteModal, setShowDeleteModal] = useState(false);

  const { data: profile } = useGetRequest({
    url: 'auth/profile',
    queryKey: 'profile'
  });

  const deleteAccount = useDeleteRequest({
    queryKey: 'profile',
    url: 'auth/account'
  });

  const handleDeleteAccount = async () => {
    try {
      await deleteAccount.mutateAsync();
      setShowDeleteModal(false);
      notifySuccess('Account deleted successfully');
      window.location.href = '/';
    } catch (err: any) {
      notifyError(err);
      setShowDeleteModal(false);
    }
  };

  const { email, first_name, last_name } = profile || {};

  const {
    handleSubmit,
    register,
    reset,
    formState: {}
  } = useForm<Inputs>({
    defaultValues: {
      first_name: first_name,
      last_name: last_name,
      email: email
    },
    resolver: zodResolver(EditProfileSchema)
  });

  const { mutateAsync } = usePatchRequest({
    url: `auth/profile`,
    queryKey: 'profile'
  });

  const onSubmit = async (data: any) => {
    const hasChanges = data.first_name !== first_name || data.last_name !== last_name;

    if (!hasChanges) {
      notifySuccess('No changes to save');
      return;
    }

    setLoading(true);

    try {
      await mutateAsync(data, {
        onSuccess: () => {
          notifySuccess('Profile updated successfully.');
        }
      });
      reset();
    } catch (err: any) {
      notifyError('Failed to update user profile details');
    }
    setLoading(false);
  };

  return (
    <DefaultLayout>
      {showDeleteModal && (
        <DeleteModal
          handleShow={() => setShowDeleteModal(false)}
          handleClick={handleDeleteAccount}
          title="Delete Your Account"
          description="Are you sure you want to delete your account? This action will permanently delete your account and all associated data. This action cannot be undone."
        />
      )}

      <div className="mx-auto max-w-270">
        <div className="grid grid-cols-5 gap-8">
          <div className="col-span-5 xl:col-span-3">
            <div className="rounded-sm border border-stroke bg-white">
              <div className="border-b border-stroke px-7 py-4">
                <h3 className="font-medium text-black">Personal Information</h3>
              </div>
              <div className="p-7">
                <form onSubmit={handleSubmit(onSubmit)}>
                  <div className="mb-5.5 flex flex-col gap-5.5 sm:flex-row">
                    <div className="w-full sm:w-1/2">
                      <Input
                        label="First Name"
                        type="text"
                        id="first_name"
                        {...register('first_name')}
                        placeholder="Enter first name"
                        defaultValue={first_name}
                        required
                      />
                    </div>

                    <div className="w-full sm:w-1/2">
                      <Input
                        label="Last Name"
                        {...register('last_name')}
                        name="last_name"
                        type="text"
                        id="last_name"
                        placeholder="Enter last name"
                        defaultValue={last_name}
                        required
                      />
                    </div>
                  </div>

                  <div className="mb-5.5">
                    <Input
                      label="Email Address"
                      name="email"
                      id="email"
                      disabled
                      type="email"
                      defaultValue={email}
                      placeholder="Enter email address"
                      readOnly
                      className="cursor-not-allowed bg-gray-100"
                    />
                  </div>

                  <div className="flex justify-end gap-4.5">
                    <Button variant="outline" type="button" onClick={() => reset()}>
                      Cancel
                    </Button>
                    <Button loading={loading} type="submit">
                      Save
                    </Button>
                  </div>
                </form>
              </div>
            </div>
          </div>
          <div className="col-span-5 xl:col-span-3">
            <div className="mt-7.5 rounded-sm border border-stroke bg-white">
              <div className="border-b border-stroke px-7 py-4">
                <h3 className="font-medium text-black">Account Settings</h3>
              </div>
              <div className="p-7">
                <div className="mb-5">
                  <h4 className="mb-2 text-danger">Delete Account</h4>
                  <p className="mb-4 text-sm text-gray-600">
                    This action will permanently delete your account and all associated data. This action cannot be
                    reversed.
                  </p>
                  <Button variant="danger" type="button" onClick={() => setShowDeleteModal(true)}>
                    Delete Account
                  </Button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </DefaultLayout>
  );
};

export default Settings;
