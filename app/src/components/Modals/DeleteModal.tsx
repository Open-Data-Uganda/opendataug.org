import { XMarkIcon } from '@heroicons/react/24/solid';
import { FC } from 'react';

interface ModalProps {
  handleShow: (args: any) => void;
  handleClick: (args: any) => void;
  title: string;
  description?: string;
}

export const DeleteModal: FC<ModalProps> = ({ handleShow, title, handleClick, description }) => {
  return (
    <div className="fixed left-0 top-[10%] z-10 my-auto w-full overflow-y-auto md:top-0">
      <div className="my-auto flex items-center justify-center px-4 pb-20 pt-4 text-center sm:p-0">
        <span className="hidden sm:inline-block sm:h-screen sm:align-middle">&#8203;</span>

        <div
          className="align-center inline-block transform overflow-hidden rounded-md border bg-white py-8 text-left outline-none transition-all focus-within:outline-none focus:outline-none sm:w-full sm:max-w-2xl sm:align-middle"
          role="dialog">
          <div className="mb-6 px-8">
            <XMarkIcon className="float-right h-6 w-6 cursor-pointer text-primary" onClick={handleShow} />
          </div>

          <h2 className="mb-4 px-8 text-center text-base font-semibold text-black">{title}</h2>

          <div className="px-8 sm:pb-6">
            <div className="relative rounded bg-white text-center">
              <p className="mb-6 text-sm leading-relaxed text-gray-600">{description || 'Are you sure?'}</p>
              <div className="flex items-center justify-center space-x-8">
                <button
                  onClick={handleShow}
                  data-modal-toggle="deleteModal"
                  type="button"
                  className="rounded border border-gray-200 bg-white px-6 py-2.5 text-sm font-medium text-gray-500 hover:bg-gray-100 focus:outline-none focus:ring-0">
                  Cancel
                </button>
                <button
                  type="submit"
                  onClick={handleClick}
                  className="rounded bg-red-600 px-6 py-2.5 text-sm font-medium text-white hover:bg-red-700 focus:outline-none focus:ring-0">
                  Confirm
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};
