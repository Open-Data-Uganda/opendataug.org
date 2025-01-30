import { FC } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';

interface PaginationProps {
  totalRecords: number;
  limit: number;
  invalidateCache?: () => void;
}

const Pagination: FC<PaginationProps> = ({ totalRecords, limit, invalidateCache }) => {
  const location = useLocation();
  const navigate = useNavigate();

  const getQueryParam = (param: string) => {
    const searchParams = new URLSearchParams(location.search);
    return searchParams.get(param);
  };

  const page = parseInt(getQueryParam('page') || '1', 10);

  const totalPages = Math.ceil(totalRecords / limit);

  const setPage = (newPage: number) => {
    const searchParams = new URLSearchParams(location.search);
    searchParams.set('page', newPage.toString());
    navigate({ search: searchParams.toString() });

    if (invalidateCache) {
      invalidateCache();
    }
  };

  return (
    <nav className="flex items-center justify-center">
      <ul className="flex h-10 items-center -space-x-px text-base">
        <li className=" cursor-pointer">
          <span
            onClick={() => setPage(Math.max(page - 1, 1))}
            className="ms-0 flex h-10 cursor-pointer items-center justify-center rounded-s-lg border border-e-0 border-gray-300 bg-white px-4 leading-tight text-gray-500 hover:bg-gray-100 hover:text-gray-700 dark:border-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">
            <span className="sr-only">Previous</span>
            <svg
              className="h-3 w-3 rtl:rotate-180"
              aria-hidden="true"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 6 10">
              <path
                stroke="currentColor"
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="2"
                d="M5 1 1 5l4 4"
              />
            </svg>
          </span>
        </li>

        {Array.from({ length: totalPages }, (_, index) => (
          <li key={index + 1} className=" cursor-pointer">
            <span
              onClick={() => setPage(index + 1)}
              className={`flex h-10 items-center justify-center border border-gray-300 bg-white px-4 leading-tight text-gray-500 hover:bg-gray-100 hover:text-gray-700  ${page === index + 1 ? 'font-bold' : ''}`}>
              {index + 1}{' '}
            </span>
          </li>
        ))}
        <li className=" cursor-pointer">
          <span
            onClick={() => setPage(Math.min(page + 1, totalPages))}
            className="flex h-10 cursor-pointer items-center justify-center rounded-e-lg border border-gray-300 bg-white px-4 leading-tight text-gray-500 hover:bg-gray-100 hover:text-gray-700 dark:border-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">
            <span className="sr-only">Next</span>
            <svg
              className="h-3 w-3 rtl:rotate-180"
              aria-hidden="true"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 6 10">
              <path
                stroke="currentColor"
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="2"
                d="m1 9 4-4-4-4"
              />
            </svg>
          </span>
        </li>
      </ul>
    </nav>
  );
};

export default Pagination;
