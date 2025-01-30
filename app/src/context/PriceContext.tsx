import React, { createContext, ReactNode, useContext, useState } from 'react';

interface PriceContextProps {
  price: number;
  setPrice: React.Dispatch<React.SetStateAction<number>>;
}

const PriceContext = createContext<PriceContextProps | undefined>(undefined);

interface PriceProviderProps {
  children: ReactNode;
}

export const PriceProvider: React.FC<PriceProviderProps> = ({ children }) => {
  const [price, setPrice] = useState(0);

  return <PriceContext.Provider value={{ price, setPrice }}>{children}</PriceContext.Provider>;
};

export const usePrice = (): PriceContextProps => {
  const context = useContext(PriceContext);
  if (context === undefined) {
    throw new Error('usePrice must be used within a PriceProvider');
  }
  return context;
};
