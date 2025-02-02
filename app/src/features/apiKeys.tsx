import { useState } from 'react';

import { ClipboardDocumentCheckIcon, ClipboardDocumentIcon, EyeIcon, EyeSlashIcon } from '@heroicons/react/24/outline';
import { useQueryClient } from '@tanstack/react-query';
import Actions from '../components/Actions';
import CreateButton from '../components/Buttons/CreateButton';
import Container from '../components/Container';
import { TableLoader } from '../components/Loaders/TableLoader';
import { DeleteModal } from '../components/Modals/DeleteModal';
import { TableData, TableError, TableHeader, TableNoData } from '../components/Tables';
import TableContainer from '../components/Tables/TableContainer';
import { notifyError, notifySuccess } from '../components/toasts';
import useDeleteRequest from '../hooks/useDeleteRequest';
import useGetRequest from '../hooks/useGetRequest';
import usePostRequest from '../hooks/usePostRequest';
import DefaultLayout from '../layout/DefaultLayout';
import { Quotation } from '../types';

const Overview: React.FC = () => {
  const { data, isError, isLoading } = useGetRequest({
    url: 'api-keys',
    queryKey: ['api-keys']
  });
  const [showModal, setShowModal] = useState(false);
  const [selected, setSelected] = useState('');
  const deleteAPIKey = useDeleteRequest('api-keys', `api-keys/${selected}`);
  const [copiedId, setCopiedId] = useState<string | null>(null);
  const [visibleId, setVisibleId] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  const handleDelete = async () => {
    try {
      await deleteAPIKey.mutateAsync();
      setShowModal(false);
      notifySuccess('API Key deleted');
    } catch (err) {
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
    } catch (err) {
      notifyError('Failed to copy API Key');
    }
  };

  const toggleVisibility = (id: string) => {
    setVisibleId(visibleId === id ? null : id);
  };

  const getHiddenValue = (value: string) => '•'.repeat(value.length);

  const { mutateAsync } = usePostRequest({
    url: 'api-keys',
    queryKey: 'api-keys'
  });

  const queryClient = useQueryClient();

  const handleCreateAPIKey = async (name: any) => {
    setLoading(true);
    try {
      await mutateAsync(name);
      notifySuccess('API Key created successfully');
    } catch (error) {
      notifyError('Error occurred while creating the API Key');
    } finally {
      setLoading(false);
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
          <CreateButton onCreateKey={handleCreateAPIKey} />
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

              <tbody>
                {isLoading && <TableLoader />}
                {isError && <TableError />}
                {data?.data === null && <TableNoData />}

                {data?.data?.map((quotation: Quotation) => (
                  <tr key={quotation.number}>
                    <TableData> Deploy</TableData>
                    <TableData>
                      <div className="flex items-center gap-2">
                        <button
                          onClick={() => toggleVisibility('9e6d336a-2f98-43e9-9fa6-11b4d5cdaf62')}
                          className="flex cursor-pointer items-center gap-2 hover:text-primary"
                          title={
                            visibleId === '9e6d336a-2f98-43e9-9fa6-11b4d5cdaf62' ? 'Hide API Key' : 'Show API Key'
                          }>
                          <span>
                            {visibleId === '9e6d336a-2f98-43e9-9fa6-11b4d5cdaf62'
                              ? '9e6d336a-2f98-43e9-9fa6-11b4d5cdaf62'
                              : getHiddenValue('9e6d336a-2f98-43e9-9fa6-11b4d5cdaf62')}
                          </span>
                          {visibleId === '9e6d336a-2f98-43e9-9fa6-11b4d5cdaf62' ? (
                            <EyeSlashIcon className="h-5 w-5" />
                          ) : (
                            <EyeIcon className="h-5 w-5" />
                          )}
                        </button>
                        {visibleId === '9e6d336a-2f98-43e9-9fa6-11b4d5cdaf62' && (
                          <button
                            onClick={() => handleCopyClick('9e6d336a-2f98-43e9-9fa6-11b4d5cdaf62')}
                            className="flex cursor-pointer items-center hover:text-primary"
                            title="Click to copy">
                            {copiedId === '9e6d336a-2f98-43e9-9fa6-11b4d5cdaf62' ? (
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
                          setShowModal(!showModal);
                          setSelected(quotation.number);
                        }}
                        detailsUrl={`quotations/${quotation.number}`}
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
