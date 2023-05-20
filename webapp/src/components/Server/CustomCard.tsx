// src/components/CustomCard.tsx
import React, { useState } from "react";

interface CustomCardProps {
  data: any;
  edit?: boolean;
}

const CustomCard: React.FC<CustomCardProps> = ({ data }) => {
  const headers = Object.keys(data);
  const [selectedHeader, setSelectedHeader] = useState(headers[0]);

  const handleSelectChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    setSelectedHeader(event.target.value);
  };

  return (
    <div className="w-full bg-gradient-to-r from-black via-gray-800 to-black p-6 rounded-lg shadow-xl shadow-blue-400/30 relative">
      <div className="overflow-x-auto">
        <select
          value={selectedHeader}
          onChange={handleSelectChange}
          className="text-white bg-gray-700 rounded-md p-2 focus:outline-none focus:ring-2 focus:ring-blue-400 focus:ring-opacity-50"
        >
          {headers.map((header, index) => (
            <option key={index} value={header}>
              {header}
            </option>
          ))}
        </select>
        <div className="mt-4 text-white">
          <span className="block text-sm font-semibold tracking-wide uppercase mb-1 text-blue-400">{selectedHeader}</span>
          <span className="block text-xl font-medium text-white">{data[selectedHeader]}</span>
        </div>
      </div>
    </div>
  );
};

export default CustomCard;
