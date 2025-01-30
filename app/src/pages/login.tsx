import { Link } from 'react-router-dom';

export default function Login() {
  return (
    <>
      <div className="border-b border-gray-200">
        <div className="mx-auto flex w-full max-w-screen-2xl items-center justify-between px-6 py-4">
          <img
            src="https://tailwindui.com/plus/img/logos/mark.svg?color=indigo&shade=900"
            alt="Uganda Data Logo"
            className="h-8"
          />
          <div className="flex items-center gap-4">
            <Link to="/login" className="text-sm font-semibold text-gray-900">
              Log in
            </Link>
            <Link
              to="/signup"
              className="rounded-lg bg-primary px-4 py-2 text-sm font-semibold text-white transition-colors hover:bg-primary/90">
              Sign up
            </Link>
          </div>
        </div>
      </div>
      {/* Rest of your login content */}
    </>
  );
}
