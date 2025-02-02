import { zodResolver } from '@hookform/resolvers/zod';
import { useState } from 'react';
import { useForm } from 'react-hook-form';
import { z } from 'zod';

import { APIKeySchema } from '../../types/schemas';
import Button from '../Button';
import Input from '../Input';
import Modal from '../Modal';

type APIKeyFormData = z.infer<typeof APIKeySchema>;

interface CreateButtonProps {
  onCreateKey?: (name: string) => Promise<void>;
}

const CreateButton = ({ onCreateKey }: CreateButtonProps) => {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [loading, setLoading] = useState(false);

  const {
    register,
    handleSubmit,
    reset,
    watch,
    formState: { errors }
  } = useForm<APIKeyFormData>({
    resolver: zodResolver(APIKeySchema)
  });

  const nameValue = watch('name');

  const onSubmit = async (data: APIKeyFormData) => {
    try {
      setLoading(true);
      if (onCreateKey) {
        await onCreateKey(data.name);
      }
      setIsModalOpen(false);
      reset();
    } catch (error) {
      console.error('Error creating API key:', error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <>
      <Button onClick={() => setIsModalOpen(true)}>Create New API Key</Button>

      <Modal isOpen={isModalOpen} onClose={() => setIsModalOpen(false)} title="Create New API Key">
        <div className="mt-4">
          <form onSubmit={handleSubmit(onSubmit)}>
            <Input
              label="API Key Name"
              placeholder="Enter a name for your API key"
              required
              maxLength={10}
              showCharCount
              value={nameValue}
              error={errors.name?.message}
              {...register('name')}
            />

            <div className="mt-6 flex justify-end space-x-3">
              <Button variant="outline" onClick={() => setIsModalOpen(false)} type="button">
                Cancel
              </Button>
              <Button type="submit" loading={loading}>
                Create
              </Button>
            </div>
          </form>
        </div>
      </Modal>
    </>
  );
};

export default CreateButton;
