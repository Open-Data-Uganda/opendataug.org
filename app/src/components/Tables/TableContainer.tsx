import React, { FC } from 'react';

interface TableProps {
  children: React.ReactNode;
}

const TableContainer: FC<TableProps> = ({ children }) => {
  return (
    <div className="w-full max-w-full  bg-white">
      <div className="bg-white pb-2.5 pt-3 ">
        <div className="max-w-full overflow-x-auto">
          <table className="w-full table-auto overflow-scroll overflow-x-auto  whitespace-nowrap text-13">
            {children}
          </table>
        </div>
      </div>
    </div>
  );
};

export default TableContainer;
