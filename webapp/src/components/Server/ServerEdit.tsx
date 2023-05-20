import React, { useState, useEffect } from 'react';
import { updateServer, getServerInfo } from '../../modules/api';

interface ServerEditProps {}

export const ServerEdit: React.FC<ServerEditProps> = () => {
  const [serverData, setServerData] = useState<any>(null);
  const [address, setAddress] = useState('');
  const [listenPort, setListenPort] = useState('');
  const [dns, setDns] = useState('');
  const [persistentKeepalive, setPersistentKeepalive] = useState('');
  const [confirmationMessage, setConfirmationMessage] = useState<string | null>(null);
  const [submitting, setSubmitting] = useState<boolean>(false);


  useEffect(() => {
    const fetchData = async () => {
      const result = await getServerInfo();
      if (result && result.status === 200 && result.server) {
        setServerData(result.server);
        setAddress(result.server.Address[0]);
        setListenPort(result.server.ListenPort.toString());
        setDns(result.server.DNS.join(', '));
        setPersistentKeepalive(result.server.PersistentKeepalive.toString());
      }
    };

    fetchData();
  }, []);

  const handleEditServer = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
  
    if (!serverData) return;
  
    setSubmitting(true);
    setConfirmationMessage(null);
  
    const updatedData = {
      ...serverData,
      Address: address ? [address] : serverData.Address,
      ListenPort: listenPort ? parseInt(listenPort) : serverData.ListenPort,
      DNS: dns ? dns.split(',').map((ip: string) => ip.trim()) : serverData.DNS,
      PersistentKeepalive: persistentKeepalive ? parseInt(persistentKeepalive) : serverData.PersistentKeepalive,
    };
  
    try {
      await updateServer(updatedData);
      setConfirmationMessage('Server configuration updated successfully.');
    } catch (error) {
      setConfirmationMessage('Error updating server configuration. Please try again.');
    } finally {
      setSubmitting(false);
    }
  };  

  const reload = () => {
    window.location.reload();
  };

  return (

    <div>
      <label htmlFor="server-edit" className=''>Edit</label>
      <input type="checkbox" id="server-edit" className="modal-toggle" />
      <div className="modal">
        <div className="modal-box relative w-full bg-gray-900 rounded-md bg-clip-padding backdrop-filter backdrop-blur-sm bg-opacity-80 border border-gray-100">
          <label htmlFor="server-edit" className="btn btn-sm btn-circle absolute right-2 top-2" onClick={reload}>✕</label>
          <h3 className="text-3xl text-blue-200 font-bold">Edit Server Configuration</h3>
          <div>
            <form onSubmit={handleEditServer}>
              <div className="space-y-4">
                <div>
                  <label htmlFor="address" className="block text-blue-100 text-md text-left my-3 font-semibold">
                    Address
                  </label>
                  <input
                    id="address"
                    value={address}
                    onChange={(e) => setAddress(e.target.value)}
                    className="input input-bordered w-full max-w-x text-black"
                  />
                </div>
                <div>
                  <label htmlFor="listenPort" className="block text-blue-100 text-md text-left my-3 font-semibold">
                    Listen Port
                  </label>
                  <input
                    id="listenPort"
                    value={listenPort}
                    onChange={(e) => setListenPort(e.target.value)}
                    className="input input-bordered w-full max-w-x text-black"
                    type="number"
                  />
                </div>
                <div>
                  <label htmlFor="dns" className="block text-blue-100 text-md text-left my-3 font-semibold">
                    DNS (comma-separated)
                  </label>
                  <input
                    id="dns"
                    value={dns}
                    onChange={(e) => setDns(e.target.value)}
                    className="input input-bordered w-full max-w-x text-black"
                  />
                </div>
                <div>
                  <label htmlFor="persistentKeepalive" className="block text-blue-100 text-md text-left my-3 font-semibold">
                    Persistent Keepalive
                  </label>
                  <input
                    id="persistentKeepalive"
                    value={persistentKeepalive}
                    onChange={(e) => setPersistentKeepalive(e.target.value)}
                    className="input input-bordered w-full max-w-x text-black"
                    type="number"
                  />
                </div>
              </div>
              <button type="submit" className="bg-gradient-to-r from-blue-300 to-blue-500 text-gray-900 font-semibold rounded-lg p-3 px-5 mt-4" disabled={submitting}>
                {submitting ? (
                  <span className="flex items-center justify-center">
                    <span className="animate-spin inline-block h-4 w-4 border-t-2 border-blue-100 rounded-full" />
                  </span>
                ) : (
                  'Submit'
                )}
              </button>
              {confirmationMessage && <p className="mt-3 text-blue-200">{confirmationMessage}</p>}
            </form>
          </div>
        </div>
      </div>
    </div>
    );
};

export default ServerEdit;
