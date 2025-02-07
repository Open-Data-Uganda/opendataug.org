import Logo from '../assets/logo.png';

const SmallHeader = () => {
  return (
    <div className="border-b border-gray-200">
      <div className="mx-auto flex w-full max-w-screen-2xl items-center justify-between px-6 py-4">
        <img src={Logo} alt="Open Data Uganda Logo" className="h-8 lg:h-10" />
      </div>
    </div>
  );
};

export default SmallHeader;
