import React, { useState } from "react";

interface CustomTableProps {
  data: any;
  edit?: boolean;
}

const CustomTable: React.FC<CustomTableProps> = ({ data }) => {
  const headers = Object.keys(data);
  const [selectedHeader, setSelectedHeader] = useState(headers[0]);

  const handleSelectChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    setSelectedHeader(event.target.value);
  };

  return (
    <div className="w-full bg-gradient-to-r from-black via-gray-800 to-black rounded-lg shadow-xl shadow-blue-400/30 relative">
      <table className="w-full text-base text-left text-white">
        {/* <thead className="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
            <tr>
              <th scope="col" className="px-6 py-3">
                Song
              </th>
              <th scope="col" className="px-6 py-3">
                Artist
              </th>
            </tr>
          </thead> */}
        <tbody>
          {headers.map(function (header, i) {
            return (
              <tr className="">
                <td className="px-6 py-4 w-80">{header}</td>
                <td className="px-6 py-4">{data[header]}</td>
              </tr>
            );
          })}
        </tbody>
      </table>
    </div>
  );
};

export default CustomTable;
