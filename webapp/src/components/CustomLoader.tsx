// src/components/CustomLoader.tsx
import React from "react";

const CustomLoader: React.FC = () => {
  return (
    <div className="flex justify-center items-center min-h-screen">
      <div className="animate-spin rounded-full h-10 w-10 border-t-2 border-b-2 border-blue-500"></div>
    </div>
  );
};

export default CustomLoader;
