import React, { FC } from 'react';

interface TableHeaderProps {
  children: React.ReactNode;
  width?: string;
}

const TableHeader: React.FC<TableHeaderProps> = ({ children, width = '220' }) => {
  return <th className={`max-w-${width}px px-4 py-4 font-bold text-black`}>{children}</th>;
};

interface TableRowProps {
  children: React.ReactNode;
}

const TableRow: React.FC<TableRowProps> = ({ children }) => {
  return <tr className="bg-[#d6d5d5] text-left dark:bg-meta-4">{children}</tr>;
};

interface TableDataProps {
  children: React.ReactNode;
}

const TableData: React.FC<TableDataProps> = ({ children }) => {
  return <td className="overflow-ellipsis px-4 py-3 text-sm text-black">{children}</td>;
};

const TableError: FC = () => {
  return (
    <tr className=" w-full">
      <TableData>Failed to fetch data</TableData>
    </tr>
  );
};

const TableNoData: FC = () => {
  return (
    <tr className=" w-full">
      <TableData>You have't created an API Key yet </TableData>
    </tr>
  );
};

export { TableData, TableError, TableHeader, TableNoData, TableRow };
