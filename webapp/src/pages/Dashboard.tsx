import React, { useContext, useEffect, useState } from "react";
import { ClientContainer } from "../components/Dashboard/clientContainer";
import { Client } from "../components/Dashboard/types";
import { getClients } from "../modules/api";
import { useAccount } from "wagmi";

import Header from "../components/Header";
import Footer from "../components/Footer";
import CustomLoader from "../components/CustomLoader";
import { AuthContext } from "../context/AuthContext";
import NotSigned from "../components/NotSigned";
import NotAuthorized from "../components/NotAuthorized";
import ServerPage from "./Server";
import DashboardLoader from "../components/DashboardLoader";

const NotConnected: React.FC = () => {
  return (
    <div className="flex h-screen flex-col">
      <Header />
      <div className="flex flex-col items-center justify-center h-screen">
        <div className="text-4xl md:text-6xl font-bold text-blue-200 mb-4">
          Not Connected
        </div>
        <div className="text-xl md:text-2xl text-white">
          Please connect your wallet to access the Dashboard
        </div>
      </div>
    </div>
  );
};

const Dashboard: React.FC = () => {
  const authContext = useContext(AuthContext);

  const [clients, setClients] = useState<Client[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const { address, isConnecting, isDisconnected } = useAccount();
  const [clientOpen, setClientOpen] = useState<boolean>(true);

  useEffect(() => {
    async function fetchClients() {
      const token = localStorage.getItem("token");
      const clientData = await getClients(token);
      setClients(clientData.clients);
      setIsLoading(false);
    }
    if (localStorage.getItem("token")) {
      fetchClients();
    }
  }, []);
  function ServerClick() {
    setClientOpen(false);
  }
  function Clientclick() {
    setClientOpen(true);
  }
  if (!address || isDisconnected) {
    return <NotConnected />;
  }

  if (isLoading || isConnecting) {
    return <DashboardLoader />;
  }

  if (!authContext?.isAuthorized) {
    return <NotAuthorized />;
  }

  if (!authContext?.isSignedIn) {
    return (
      <div>
        <NotSigned component="Dashboard" />
      </div>
    );
  }

  return (
    <div>
      <Header />
      <div className="flex flex-row justify-center h-16 mx-4 px-64 mt-4 mb-8">
        <div className="flex flex-row border border-blue-200 rounded-xl shadow-lg shadow-blue-300/30 bg-gray-800 bg-opacity-40">
          <div className="flex justify-center">
            <button
              onClick={Clientclick}
              className="px-5 rounded-s-lg text-blue-200 border-blue hover:bg-blue-300 hover:border-black hover:text-black font-bold transition focus:ring focus:ring-blue-500 focus:ring-opacity-80"
            >
              Client
            </button>
          </div>
          <div className="flex justify-center">
            <button
              onClick={ServerClick}
              className="px-5 rounded-e-lg text-blue-200 border-blue hover:bg-blue-300 hover:border-black hover:text-black font-bold transition focus:ring focus:ring-blue-500 focus:ring-opacity-80"
            >
              Server
            </button>
          </div>
        </div>
      </div>
      {clientOpen && <ClientContainer clients={clients} />}
      {!clientOpen && <ServerPage />}
      <Footer />
    </div>
  );
};

export default Dashboard;
