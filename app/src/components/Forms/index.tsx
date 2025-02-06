import { ChevronDownIcon } from '@heroicons/react/24/outline';
import { FC, ReactNode } from 'react';

interface LabelProps {
  text: string;
  id: string;
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

export { Label, QuotationLabel, Select };
