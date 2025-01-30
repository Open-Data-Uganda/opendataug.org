import { EyeIcon, PencilSquareIcon, TrashIcon } from '@heroicons/react/24/outline';
import React, { useEffect, useRef, useState } from 'react';
import { Link } from 'react-router-dom';

interface ActionsProps {
  detailsUrl: string;
  onTrashClick: () => void;
}

const Actions: React.FC<ActionsProps> = ({ detailsUrl, onTrashClick }) => {
  const useClickOutside = (handler: () => void) => {
    const domNode = useRef<HTMLDivElement>(null);

    useEffect(() => {
      const maybeHandler = (event: MouseEvent) => {
        if (domNode.current && !domNode.current.contains(event.target as Node)) {
          handler();
        }
      };

      document.addEventListener('mousedown', maybeHandler);

      return () => {
        document.removeEventListener('mousedown', maybeHandler);
      };
    }, [handler]);

    return domNode;
  };

  const [dropdownOpen, setDropdownOpen] = useState(false);

  const domNode = useClickOutside(() => {
    setDropdownOpen(false);
  });

  return (
    <div className="  text-center" ref={domNode}>
      <button
        onClick={() => setDropdownOpen(!dropdownOpen)}
        className="flex items-center justify-center text-xl font-bold text-black">
        ...
      </button>

      <div className={` ${dropdownOpen ? 'relative z-10' : 'hidden'}  divide-y divide-gray-100 rounded `}>
        <div className="absolute -left-10 z-10 w-auto justify-start rounded-md border border-[#eee] bg-white py-2 text-13 text-gray-700 shadow ">
          <Link className="flex flex-row items-center px-3 py-2" to={detailsUrl}>
            <span>
              <EyeIcon className="mr-3 h-4 text-sky-700" />
            </span>
            <span className="pr-3 text-sm">View</span>
          </Link>

          <Link className=" flex flex-row items-center px-3 py-2" to={''}>
            <span>
              <PencilSquareIcon className="mr-3 h-4" />
            </span>
            <span className=" pr-3 text-sm">Edit</span>
          </Link>

          <span className="flex cursor-pointer flex-row items-center px-3 py-2" onClick={onTrashClick}>
            <button className="  text-red-700 hover:text-primary">
              <TrashIcon className="mr-3 h-4" />
            </button>
            <span className=" pr-3 text-sm">Delete</span>
          </span>
        </div>
      </div>
    </div>
  );
};

export default Actions;
