import { useState } from 'react';

import { ClipboardDocumentCheckIcon, ClipboardDocumentIcon, EyeIcon, EyeSlashIcon } from '@heroicons/react/24/outline';
import { useQueryClient } from '@tanstack/react-query';
import Actions from '../components/Actions';
import Container from '../components/Container';
import { TableLoader } from '../components/Loaders/TableLoader';
import { DeleteModal } from '../components/Modals/DeleteModal';
import { TableData, TableError, TableHeader, TableNoData } from '../components/Tables';
import TableContainer from '../components/Tables/TableContainer';
import { notifyError, notifySuccess } from '../components/toasts';
import { useDeleteRequest } from '../hooks/useDeleteRequest';
import { useGetRequest } from '../hooks/useGetRequest';
import { usePostRequest } from '../hooks/usePostRequest';
import DefaultLayout from '../layout/DefaultLayout';
import { APIKey } from '../types';
import CreateAPIKey from './CreateAPIKey';

const Overview: React.FC = () => {
  const { data, isError, isLoading } = useGetRequest({
    url: 'api-keys',
    queryKey: ['api-keys'],
  });

  const [showModal, setShowModal] = useState(false);
  const [selected, setSelected] = useState('');
  const deleteAPIKeyMutation = useDeleteRequest({
    queryKey: 'api-keys',
    url: `api-keys/${selected}`,
  });

  const [copiedId, setCopiedId] = useState<string | null>(null);
  const [visibleId, setVisibleId] = useState<string | null>(null);

  const handleDelete = async () => {
    try {
      await deleteAPIKeyMutation.mutateAsync();
      setShowModal(false);
      notifySuccess('API Key deleted');
    } catch {
      notifyError('API Key not deleted');
      setShowModal(false);
    }
  };

  const handleCopyClick = async (text: string) => {
    try {
      await navigator.clipboard.writeText(text);
      setCopiedId(text);
      notifySuccess('API Key copied to clipboard');
      setTimeout(() => setCopiedId(null), 2000);
    } catch {
      notifyError('Failed to copy API Key');
    }
  };

  const toggleVisibility = (id: string) => {
    setVisibleId(visibleId === id ? null : id);
  };

  const getHiddenValue = (value: string) => '•'.repeat(value.length);

  const { mutateAsync } = usePostRequest({
    url: 'api-keys',
    queryKey: 'api-keys',
  });

  const queryClient = useQueryClient();

  const handleCreateAPIKey = async (data: any) => {
    try {
      await mutateAsync(data);
      notifySuccess('API Key created successfully');
    } catch (error: any) {
      if (error?.response?.status === 409) {
        notifyError('An API key with this name already exists');
      } else if (error?.response?.status === 400) {
        notifyError('API Key name is required');
      } else {
        notifyError('Error occurred while creating the API Key');
      }
    } finally {
      queryClient.invalidateQueries({ queryKey: ['api-keys'] });
    }
  };

  return (
    <DefaultLayout>
      {showModal && (
        <DeleteModal handleClick={handleDelete} title="Delete API Key" handleShow={() => setShowModal(!showModal)} />
      )}

      <Container>
        <div className="flex flex-row justify-end">
          <CreateAPIKey onCreateKey={handleCreateAPIKey} />
        </div>

        <div className="mt-4 grid grid-cols-12 gap-4 md:mt-6 md:gap-6 2xl:mt-7.5 2xl:gap-7.5">
          <div className="col-span-12 mb-10">
            <TableContainer>
              <thead>
                <tr className="bg-gray-2 text-left">
                  <TableHeader> Name</TableHeader>
                  <TableHeader> API Key</TableHeader>
                  <TableHeader width="0">Actions</TableHeader>
                </tr>
              </thead>

              <tbody className=" h-20">
                {isLoading && <TableLoader />}
                {isError && <TableError />}
                {data?.length === 0 && <TableNoData />}

                {data?.map((api_key: APIKey) => (
                  <tr key={api_key.id}>
                    <TableData> {api_key.name}</TableData>
                    <TableData>
                      <div className="flex items-center gap-2">
                        <button
                          onClick={() => toggleVisibility(api_key.key)}
                          className="flex cursor-pointer items-center gap-2 hover:text-primary"
                          title={visibleId === api_key.key ? 'Hide API Key' : 'Show API Key'}>
                          <span>{visibleId === api_key.key ? api_key.key : getHiddenValue(api_key.key)}</span>
                          {visibleId === api_key.key ? (
                            <EyeSlashIcon className="h-5 w-5" />
                          ) : (
                            <EyeIcon className="h-5 w-5" />
                          )}
                        </button>
                        {visibleId === api_key.key && (
                          <button
                            onClick={() => handleCopyClick(api_key.key)}
                            className="flex cursor-pointer items-center hover:text-primary"
                            title="Click to copy">
                            {copiedId === api_key.key ? (
                              <ClipboardDocumentCheckIcon className="h-5 w-5 text-green-500" />
                            ) : (
                              <ClipboardDocumentIcon className="h-5 w-5" />
                            )}
                          </button>
                        )}
                      </div>
                    </TableData>
                    <TableData>
                      <Actions
                        onTrashClick={() => {
                          setSelected(api_key.id);
                          setShowModal(!showModal);
                        }}
                      />
                    </TableData>
                  </tr>
                ))}
              </tbody>
            </TableContainer>
          </div>
        </div>
      </Container>
    </DefaultLayout>
  );
};

export default Overview;
