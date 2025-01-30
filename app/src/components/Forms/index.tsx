import { ChevronDownIcon } from '@heroicons/react/24/outline';
import { FC, ReactNode } from 'react';
import { useFormContext } from 'react-hook-form';
import { ErrorMessage } from '../ErrorMessage';

interface LabelProps {
  text: string;
  id: string;
}

interface InputProps {
  type: string;
  placeholder?: string;
  name: string;
  validation: object;
  value?: any;
  required?: boolean;
  onChange?: (e: any) => void;
}

interface SelectProps {
  children: ReactNode;
}

const Label: FC<LabelProps> = ({ text, id }) => {
  return (
    <label className="mb-3 block text-black" id={id}>
      {text}
    </label>
  );
};

interface QuotationLabelProps {
  text?: string;
  required?: boolean;
  name: string;
}

const QuotationLabel: FC<QuotationLabelProps> = ({ name, required, text }) => {
  return (
    <label
      className="block text-sm font-semibold leading-6 text-gray-900 first-letter:capitalize"
      htmlFor={name}
      id={name}>
      {text} {required && <span className=" text-red-500">*</span>}
    </label>
  );
};

const Select: FC<SelectProps> = ({ children }) => {
  return (
    <div className="relative z-20">
      <select className="relative w-full appearance-none rounded border border-gray-400 bg-transparent px-3 py-2.5 outline-none  transition hover:border-blue-700 focus:border-primary">
        {children}
      </select>
      <span className="absolute right-4 top-1/2 z-10 -translate-y-1/2">
        <ChevronDownIcon className=" h-4" />
      </span>
    </div>
  );
};

const Input: React.FC<InputProps> = ({ type, placeholder, onChange, name, value, required, validation }) => {
  const {
    register,
    formState: { errors }
  } = useFormContext();

  return (
    <>
      <input
        {...register(name, validation)}
        className={`block w-full appearance-none rounded border px-3 py-2.5 leading-normal text-gray-600 hover:border-blue-700 focus:bg-white focus:outline-none ${
          errors[name] ? 'border-red-500' : 'border-gray-400'
        }`}
        id={name}
        required={required}
        onChange={onChange}
        value={value}
        type={type}
        placeholder={placeholder}
      />
      {errors[name]?.message && <ErrorMessage error={errors[name]?.message?.toString()} />}
    </>
  );
};

export { Input, Label, QuotationLabel, Select };
