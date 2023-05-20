import React, { useState, useEffect, useContext } from "react";
import CustomCard from "../components/Server/CustomCard";
import CustomLoader from "../components/CustomLoader";
import { getStatus, getServerInfo, getServerConfig } from "../modules/api";
import { motion } from "framer-motion";
import Header from "../components/Header";
import Footer from "../components/Footer";
import { saveAs } from "file-saver";
import ServerEdit from "../components/Server/ServerEdit";
import { useAccount } from "wagmi";
import NotSigned from "../components/NotSigned";
import { AuthContext } from "../context/AuthContext";
import NotAuthorized from "../components/NotAuthorized";
import CustomTable from "../components/Server/CustomTable";

const NotConnected: React.FC = () => {
  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <div className="text-4xl md:text-6xl font-bold text-blue-200 mb-4">
        Not Connected
      </div>
      <div className="text-xl md:text-2xl text-white">
        Please connect your wallet to access the Server
      </div>
    </div>
  );
};

const ServerPage: React.FC = () => {
  const [status, setStatus] = useState<any>(null);
  const [serverInfo, setServerInfo] = useState<any>(null);
  const [serverConfig, setServerConfig] = useState<any>(null);
  const [isLoading, setIsLoading] = useState(true);
  const { address, isConnecting, isDisconnected } = useAccount();
  const authContext = useContext(AuthContext);

  const handleDownload = async () => {
    try {
      const config = serverConfig;
      const blob = new Blob([config], { type: "text/plain;charset=utf-8" });
      saveAs(blob, `server_config.txt`);
    } catch (error) {}
  };

  useEffect(() => {
    const fetchData = async () => {
      try {
        const statusData = await getStatus();
        setStatus(statusData);
        const serverInfoData = await getServerInfo();
        setServerInfo(serverInfoData.server);
        const serverConfigData = await getServerConfig();
        setServerConfig(serverConfigData);
      } catch (error) {
      } finally {
        setIsLoading(false);
      }
    };

    fetchData();
  }, []);

  if (!address || isDisconnected) {
    return <NotConnected />;
  }

  if (isLoading || isConnecting) {
    return <CustomLoader />;
  }

  if (!authContext?.isAuthorized) {
    return <NotAuthorized />;
  }

  if (!authContext?.isSignedIn) {
    return <NotSigned component="Server" />;
  }

  return (
    <>
      <motion.div
        className="container mx-auto px-8 py-8 border border-blue-200 rounded-xl shadow-xl shadow-blue-300/30 bg-gray-800 bg-opacity-40 mt-10"
        initial={{ opacity: 0 }}
        animate={{ opacity: 1 }}
        transition={{ duration: 0.5 }}
      >
        <div className="container mx-auto px-4">
          <div className="flex justify-between">
            <button
              onClick={() => handleDownload()}
              className="bg-gradient-to-br bg-black border border-blue-300 text-white px-6 py-2 rounded-lg shadow-lg transition-all hover:from-blue-300 hover:to-blue-200 hover:text-black hover:border-black mb-10 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
            >
              Download Config
            </button>
            <div className="px-6 py-2 rounded-lg shadow-lg transition-all mb-10 border border-black bg-blue-300 text-black hover:bg-black hover:border-blue-300 hover:text-blue-200">
              <ServerEdit />
            </div>
          </div>

          <div className="mb-5">
            <h1 className="text-transparent bg-clip-text bg-gradient-to-r from-blue-400 via-gray-100 to-blue-200 font-bold text-2xl md:text-3xl mb-4 mt-8">
              Server Status
            </h1>
            {status && <CustomTable data={status} edit={false} />}
            <h1 className="text-transparent bg-clip-text bg-gradient-to-r from-blue-400 via-gray-100 to-blue-200 font-bold text-2xl md:text-3xl mb-4 mt-8">
              Server Information
            </h1>
            {serverInfo && <CustomTable data={serverInfo} edit={false} />}
          </div>
        </div>
      </motion.div>
      <Footer />
    </>
  );
};

export default ServerPage;
