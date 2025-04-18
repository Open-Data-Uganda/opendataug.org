import { ArrowRightEndOnRectangleIcon } from '@heroicons/react/24/outline';
import { useEffect, useRef, useState } from 'react';
import { Link } from 'react-router-dom';
import defaultImg from '../../assets/default.png';

import { useAuth } from '../../context/AuthContext';
import { useGetRequest } from '../../hooks/useGetRequest';

const DropdownUser = () => {
  const [dropdownOpen, setDropdownOpen] = useState(false);
  const { logout } = useAuth();

  const trigger = useRef<any>(null);
  const dropdown = useRef<any>(null);

  useEffect(() => {
    const clickHandler = ({ target }: MouseEvent) => {
      if (!dropdown.current) return;
      if (!dropdownOpen || dropdown.current.contains(target) || trigger.current.contains(target)) return;
      setDropdownOpen(false);
    };
    document.addEventListener('click', clickHandler);
    return () => document.removeEventListener('click', clickHandler);
  });

  // close if the esc key is pressed
  useEffect(() => {
    const keyHandler = ({ key }: KeyboardEvent) => {
      if (!dropdownOpen || key !== 'Escape') return;
      setDropdownOpen(false);
    };
    document.addEventListener('keydown', keyHandler);
    return () => document.removeEventListener('keydown', keyHandler);
  });

  const { data: profile, isLoading } = useGetRequest({
    url: 'auth/profile',
    queryKey: 'profile',
  });

  if (isLoading) {
    return <div>Loading...</div>;
  }

  const handleLogout = async () => {
    await logout();

    if (window.location.pathname !== '/login') {
      window.location.href = '/login';
    }
  };

  return (
    <div className="relative">
      <Link ref={trigger} onClick={() => setDropdownOpen(!dropdownOpen)} className="flex items-center gap-4" to="#">
        <span className="hidden text-right text-13   lg:block">
          <span className="block text-sm font-medium capitalize text-black dark:text-white">{profile?.name}</span>
        </span>

        <span className="h-12 w-12 rounded-full">
          <img src={defaultImg} className="rounded-full" alt="User" />
        </span>

        <svg
          className="hidden fill-current sm:block"
          width="12"
          height="8"
          viewBox="0 0 12 8"
          fill="none"
          xmlns="http://www.w3.org/2000/svg">
          <path
            fillRule="evenodd"
            clipRule="evenodd"
            d="M0.410765 0.910734C0.736202 0.585297 1.26384 0.585297 1.58928 0.910734L6.00002 5.32148L10.4108 0.910734C10.7362 0.585297 11.2638 0.585297 11.5893 0.910734C11.9147 1.23617 11.9147 1.76381 11.5893 2.08924L6.58928 7.08924C6.26384 7.41468 5.7362 7.41468 5.41077 7.08924L0.410765 2.08924C0.0853277 1.76381 0.0853277 1.23617 0.410765 0.910734Z"
            fill=""
          />
        </svg>
      </Link>

      <div
        ref={dropdown}
        onFocus={() => setDropdownOpen(true)}
        onBlur={() => setDropdownOpen(false)}
        className={`absolute right-0 mt-4 flex w-62.5 flex-col rounded-sm border border-stroke bg-white shadow-default dark:border-strokedark dark:bg-boxdark ${
          dropdownOpen === true ? 'block' : 'hidden'
        }`}>
        <button
          onClick={handleLogout}
          className="flex items-center gap-3.5 px-6 py-4 text-13 font-medium duration-300 ease-in-out hover:text-primary  ">
          <ArrowRightEndOnRectangleIcon className=" h-5" />
          Log Out
        </button>
      </div>
    </div>
  );
};

export default DropdownUser;
