import { FC } from 'react';

interface LoadingButtonProps {
  loading: boolean;
  text: string;
}

export const LoadingButton: FC<LoadingButtonProps> = ({ loading, text }) => {
  return (
    <button className="my-4 inline-flex items-center justify-center rounded bg-primary px-5 py-2 text-center font-medium text-white hover:bg-opacity-90 md:my-8 lg:px-6 xl:px-8">
      {!loading ? (
        text
      ) : (
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
      )}
    </button>
  );
};
