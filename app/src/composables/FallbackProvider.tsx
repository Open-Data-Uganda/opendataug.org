import * as React from 'react';

export type FallbackType = NonNullable<React.ReactNode> | null;

export interface FallbackContextType {
  updateFallback: (fallback: FallbackType) => void;
}

export const FallbackContext = React.createContext<FallbackContextType>({
  // eslint-disable-next-line @typescript-eslint/no-empty-function
  updateFallback: () => {}
});

interface FallbackProviderProps {
  children: React.ReactNode;
}

export const FallbackProvider: React.FC<FallbackProviderProps> = ({ children }) => {
  const [fallback, setFallback] = React.useState<FallbackType>(null);

  const updateFallback = React.useCallback((newFallback: FallbackType) => {
    setFallback(newFallback);
  }, []);

  const renderChildren = React.useMemo(() => {
    return children;
  }, [children]);

  return (
    <FallbackContext.Provider value={{ updateFallback }}>
      <React.Suspense fallback={fallback}>{renderChildren}</React.Suspense>
    </FallbackContext.Provider>
  );
};
