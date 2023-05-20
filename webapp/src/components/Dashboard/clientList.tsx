// clientList.tsx
import React, {useState} from 'react';
import { Client } from './types';
import { HiOutlineMail , HiDownload } from 'react-icons/hi';
import { MdDelete } from 'react-icons/md';
import ClientEdit from './clientEdit';
import { emailClientConfig, getClientConfig, deleteClient, updateClient, getClientInfo } from '../../modules/api';
import { saveAs } from 'file-saver';
import { IconContext } from "react-icons";

interface ClientListProps {
  clients: Client[];
}

export const ClientList: React.FC<ClientListProps> = ({ clients }) => {
  const [enabledClients, setEnabledClients] = useState<Record<string, boolean>>(
    clients.reduce((acc: Record<string, boolean>, client) => {
      acc[client.UUID] = client.Enable;
      return acc;
    }, {}),
  );

  const handleEmail = async (clientId:string) => {
    try {
      const response = await emailClientConfig(clientId)
      if (!response) {
      }
    } catch (error) {
    }
  }

  const handleDelete = async (clientId:string) => {
    try {
      const response = await deleteClient(clientId)
      if (!response) {
      }
      window.location.reload();
    } catch (error) {
    }
  }
  
  const handleDownload = async (clientId: string) => {
    try {
      const response = await getClientConfig(clientId);
      
      if (!response) {
      }
  
      const configText = response.data;
      const blob = new Blob([configText], { type: 'text/plain;charset=utf-8' });
      saveAs(blob, `${clientId}_config.conf`);
    } catch (error) {
    }
  };

  const handleClientAccess = async (clientId: string) => {
    const client = clients.find((c) => c.UUID === clientId);
    if (!client) {
      return;
    }
    try {

      const response = await getClientInfo(clientId);
      if (!response) {
      }

      const clientData = response.data.client;
      clientData.Enable = !clientData.Enable;
      
      const updateResponse = await updateClient(clientId, clientData);
      if (!updateResponse) {
      }
    } catch (error) {
    }

    setEnabledClients((prev) => {
      const newEnabledClients = { ...prev, [clientId]: !prev[clientId] };
      return newEnabledClients;
    });
  };

  return (
    <div className="w-full bg-gradient-to-r from-black via-gray-800 to-black p-6 rounded-lg shadow-xl shadow-blue-400/30">
      <div className="overflow-x-auto">
        <table className="min-w-full divide-y divide-blue-500 text-white text-left">
          <thead>
            <tr>
              <th className="px-4 py-2">Name</th>
              <th className="px-4 py-2">Email</th>
              <th className="px-4 py-2">Tags</th>
              <th className="px-4 py-2">Status</th>
              <th className="px-4 py-2">Created</th>
              <th className="px-4 py-2">Updated</th>
              <th className="px-4 py-2">Enable/Disable </th>
              <th className="px-4 py-2">Actions</th>
            </tr>
          </thead>
          <tbody>
            {clients.map((client, index) => (
              <tr key={client.UUID} className={index % 2 === 0 ? 'bg-gradient-to-r from-black via-gray-800 to-black' : 'bg-gradient-to-r from-black via-gray-900 to-black'}>
                <td className="px-4 py-4 whitespace-nowrap">
                  <div className="text-sm font-medium">{client.Name}</div>
                </td>
                <td className="px-4 py-4 whitespace-nowrap">
                  <div className="text-sm">{client.Email}</div>
                </td>
                <td className="px-4 py-4 whitespace-nowrap">
                  <div className="text-sm">{client.Tags.join(', ')}</div>
                </td>
                <td className="px-4 py-4 whitespace-nowrap">
                <span className={`px-2 inline-flex text-xs leading-5 font-semibold rounded-full ${enabledClients[client.UUID] ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'}`}>
                  {enabledClients[client.UUID] ? 'Enabled' : 'Disabled'}
                </span>
                </td>
                <td className="px-4 py-4 whitespace-nowrap">
                  <div className="text-sm">{new Date(client.Created).toLocaleDateString()}</div>
                </td>
                <td className="px-4 py-4 whitespace-nowrap">
                  <div className="text-sm">{new Date(client.Updated).toLocaleDateString()}</div>
                </td>
                <td className="px-4 py-4 whitespace-nowrap">
                  <div className="text-sm flex space-x-2">
                    <input
                      id="enable"
                      type="checkbox"
                      checked={enabledClients[client.UUID] || false}
                      onChange={() => handleClientAccess(client.UUID)}
                      className="toggle toggle-sm my-3"
                    />
                  </div>
                </td>

                <td className="px-4 py-4 whitespace-nowrap">
                  <div className="text-sm flex space-x-2">
                    <IconContext.Provider value={{ className: "h-5 w-5 cursor-pointer" }}>
                      <div onClick={() => handleEmail(client.UUID)}>
                        <HiOutlineMail />
                      </div>
                      <ClientEdit client={client}/>
                      <div onClick={() => handleDownload(client.UUID)}>
                        <HiDownload />
                      </div>
                      <div onClick={() => handleDelete(client.UUID)}>
                        <MdDelete />
                      </div>
                    </IconContext.Provider>
                  </div>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};
