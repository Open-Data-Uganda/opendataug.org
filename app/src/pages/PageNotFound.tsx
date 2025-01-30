import { useNavigate } from 'react-router-dom';

const PageNotFound = () => {
  const navigate = useNavigate();

  const goBack = () => {
    navigate(-1);
  };

  return (
    <section className="bg-white dark:bg-gray-900 ">
      <div className="container mx-auto flex min-h-screen items-center px-6 py-12">
        <div className="mx-auto flex max-w-sm flex-col items-center text-center">
          <p className="rounded-full bg-blue-50 p-3 text-sm font-medium text-blue-500 dark:bg-gray-800">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth="2"
              stroke="currentColor"
              className="h-6 w-6">
              <path
                strokeLinecap="round"
                strokeWidth="round"
                d="M12 9v3.75m9-.75a9 9 0 11-18 0 9 9 0 0118 0zm-9 3.75h.008v.008H12v-.008z"
              />
            </svg>
          </p>

          <h1 className="mt-3 text-2xl font-semibold text-gray-800 dark:text-white md:text-3xl">Page not found</h1>

          <div className="mt-6 flex w-full shrink-0 items-center gap-x-3 sm:w-auto">
            <button
              className="flex items-center justify-center gap-x-2 rounded border bg-gray-300 px-10 py-2 text-sm text-gray-700 transition-colors duration-200 hover:bg-gray-100 dark:border-gray-700 dark:bg-gray-900 dark:text-gray-200 dark:hover:bg-gray-800 sm:w-auto"
              onClick={goBack}>
              Go back
            </button>
          </div>
        </div>
      </div>
    </section>
  );
};

export default PageNotFound;
