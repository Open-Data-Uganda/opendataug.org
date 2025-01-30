const Loader = () => {
  return (
    <span className="flex h-screen w-full items-center justify-center bg-white">
      <span className="h-16 w-16 animate-spin rounded-full border-4 border-solid border-primary border-t-transparent" />
    </span>
  );
};

const TableLoader = () => {
  return (
    <tr className="flex h-screen w-full items-center justify-center bg-white">
      <td className="h-8 w-8 animate-spin rounded-full border-4 border-solid border-primary border-t-transparent"></td>
    </tr>
  );
};

export { Loader, TableLoader };
