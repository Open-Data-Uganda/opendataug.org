import { forwardRef } from 'react';
import { ErrorMessage } from './ErrorMessage';

interface InputProps extends React.InputHTMLAttributes<HTMLInputElement> {
  label: string;
  error?: string;
  required?: boolean;
  showCharCount?: boolean;
  maxLength?: number;
  value?: string;
}

const Input = forwardRef<HTMLInputElement, InputProps>(
  ({ label, error, required, className = '', showCharCount, maxLength, value = '', ...props }, ref) => {
    return (
      <div className=" mb-4">
        <div className="mb-2 flex items-center justify-between">
          <label className="block text-[12px]">
            {label} {required && <span className="text-red-500">*</span>}
          </label>
          {showCharCount && maxLength && (
            <span className="text-xs text-gray-500">
              {value.length}/{maxLength}
            </span>
          )}
        </div>
        <div className="relative">
          <input
            ref={ref}
            maxLength={maxLength}
            className={`w-full rounded border py-2 p-3 focus:border-blue-500 focus:ring-2 focus:ring-blue-500 ${className}`}
            {...props}
          />
        </div>
        {error && <ErrorMessage error={error} />}
      </div>
    );
  }
);

Input.displayName = 'Input';

export default Input;
